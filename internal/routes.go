package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ramazanyigit/ephemeris/internal/services"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")
	{
		api.Get("/status", func(c *fiber.Ctx) error {
			return c.Status(200).JSON(fiber.Map{
				"status": "ok",
				"name": "ephemeris",
				"version": "1.0.0",
			})
		})

		api.Get("/log", services.ReadDiaryLogs)
		api.Post("/log", services.CreateDiaryLog)
	}
}