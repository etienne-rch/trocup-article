package handlers

import (
	"trocup-article/models"
	"trocup-article/services"

	"github.com/gofiber/fiber/v2"
)

func UpdateArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	article := new(models.Article)
	if err := c.BodyParser(article); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	updatedArticle, err := services.UpdateArticle(id, article)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(updatedArticle)
}
