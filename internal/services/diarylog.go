package services

import (
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"github.com/ramazanyigit/ephemeris/internal/cryptography"
	"github.com/ramazanyigit/ephemeris/internal/database"
	"github.com/ramazanyigit/ephemeris/internal/model"
)

func ReadDiaryLogs(c *fiber.Ctx) error {
	db := database.GetConnection()

	var diaryLogs []*model.DiaryLog
	db.Find(&diaryLogs)

	for _, element := range diaryLogs {
		decodedEntry, _ := base64.StdEncoding.DecodeString(element.Entry)
		decryptedEntry, _ := cryptography.Decrypt(decodedEntry)
		element.Entry = string(decryptedEntry)
	}

	return c.JSON(diaryLogs)
}

func CreateDiaryLog(c *fiber.Ctx) error {
	db := database.GetConnection()

	var diaryLog model.DiaryLog
	if err := c.BodyParser(&diaryLog); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Invalid request body"})
	}

	encryptedEntry, _ := cryptography.Encrypt([]byte(diaryLog.Entry))
	diaryLog.Entry = base64.StdEncoding.EncodeToString(encryptedEntry)
	db.Create(&diaryLog)

	return c.Status(fiber.StatusCreated).JSON(diaryLog)
}