package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ramazanyigit/ephemeris/internal"
	"github.com/ramazanyigit/ephemeris/internal/cryptography"
	"github.com/ramazanyigit/ephemeris/internal/database"
	"log"
)

func main() {
	log.Println("Starting database connection...")
	database.CreateConnection()

	log.Println("Running auto-migration for models...")
	database.AutoMigration()

	log.Println("Controlling encryption key pair...")
	cryptography.CreateKeyPairIfNotExists()

	log.Print("Creating web application...")
	app := fiber.New()

	log.Println("Registering static content...")
	app.Static("/", "./public")
	log.Println("Registered static content at \"/\" from \"./public\"")

	log.Println("Registering routes...")
	internal.RegisterRoutes(app)
	log.Println("Registered routes successfully")

	log.Println("Starting application at port 3000.")
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}