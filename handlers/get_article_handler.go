package handlers

import (
	"trocup-article/services"

	"github.com/gofiber/fiber/v2"
)

func GetArticleByID(c *fiber.Ctx) error {
	id := c.Params("id")
	article, err := services.GetArticleByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Article not found"})
	}
	return c.JSON(article)
}
