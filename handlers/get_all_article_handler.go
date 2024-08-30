package handlers

import (
	"trocup-article/services"

	"github.com/gofiber/fiber/v2"
)

func GetArticles(c *fiber.Ctx) error {
	articles, err := services.GetAllArticles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(articles)
}
