package handlers

import (
	"trocup-article/services"

	"github.com/gofiber/fiber/v2"
)

func UpdateArticle(c *fiber.Ctx) error {
	articleID := c.Params("id")

	// Appel au service pour mettre à jour l'article
	updatedArticle, err := services.UpdateArticle(c, articleID)
	if err != nil {
		if err.Error() == "article not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Article not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update article"})
	}

	return c.JSON(updatedArticle)
}
