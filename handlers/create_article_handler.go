package handlers

import (
	"trocup-article/models"
	"trocup-article/services"

	"github.com/gofiber/fiber/v2"
)

func CreateArticle(c *fiber.Ctx) error {
	article := new(models.Article)
	if err := c.BodyParser(article); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	createdArticle, err := services.CreateArticle(article)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(createdArticle)
}
