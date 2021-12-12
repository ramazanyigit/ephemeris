package cryptography

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

const KEY_LOCATION = "./keys/private.pem"
const KEY_LOCATION_PUBLIC = "./keys/public.pem"
const DEFAULT_PASSWORD = "defaultPassword"

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func CreateKeyPairIfNotExists() {
	isPrivateKeyExists := true
	isPublicKeyExists := true

	if _, err := os.Stat(KEY_LOCATION); errors.Is(err, os.ErrNotExist) {
		isPrivateKeyExists = false
	}

	if _, err := os.Stat(KEY_LOCATION_PUBLIC); errors.Is(err, os.ErrNotExist) {
		isPublicKeyExists = false
	}


	if !isPrivateKeyExists || !isPublicKeyExists {
		generateKeyPair()
	} else {
		readPrivateKey()
		readPublicKey()
	}
}

func readPrivateKey() {
	pemContent, err := ioutil.ReadFile(KEY_LOCATION)

	if err != nil {
		log.Fatal(err)
	}

	block, _ := pem.Decode(pemContent)
	decryptedBlock, err := x509.DecryptPEMBlock(block, []byte(DEFAULT_PASSWORD))

	if err != nil {
		log.Fatal(err)
	}

	privateKey, err = x509.ParsePKCS1PrivateKey(decryptedBlock)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Private key successfully loaded.")
}

func readPublicKey() {
	pemContent, err := ioutil.ReadFile(KEY_LOCATION_PUBLIC)

	if err != nil {
		log.Fatal(err)
	}

	block, _ := pem.Decode(pemContent)
	parsedPublicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	publicKey = parsedPublicKey.(*rsa.PublicKey)
	log.Println("Public key successfully loaded.")
}

func generateKeyPair() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)

	if err != nil {
		panic("Cannot create the private key: " + err.Error())
	}

	privateKeyPemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// TODO: Replace with a more secure method
	privateKeyPemBlock, err = x509.EncryptPEMBlock(rand.Reader, privateKeyPemBlock.Type, privateKeyPemBlock.Bytes, []byte(DEFAULT_PASSWORD), x509.PEMCipherAES256)

	if err != nil {
		panic("Cannot encrypt the private key: " + err.Error())
	}

	privatePemFile, err := os.Create(KEY_LOCATION)

	if err != nil {
		panic("Cannot create the private key file: " + err.Error())
	}

	err = pem.Encode(privatePemFile, privateKeyPemBlock)
	if err != nil {
		panic("Cannot encode the private key: " + err.Error())
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)

	if err != nil {
		panic("Cannot marshal the public key: " + err.Error())
	}

	publicKeyPemBlock := &pem.Block{
		Type: "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	publicPemFile, err := os.Create(KEY_LOCATION_PUBLIC)

	if err != nil {
		panic("Cannot create the public key file: " + err.Error())
	}

	err = pem.Encode(publicPemFile, publicKeyPemBlock)
	if err != nil {
		panic("Cannot encode the public key: " + err.Error())
	}

	log.Println("Created key pair with default password.")
}

func Encrypt(data []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, data, nil)
}

func Decrypt(data []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, data, nil)
}
