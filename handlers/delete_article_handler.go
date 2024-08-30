package handlers

import (
	"trocup-article/services"

	"github.com/gofiber/fiber/v2"
)

func DeleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	err := services.DeleteArticle(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
