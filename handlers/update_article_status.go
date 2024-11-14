package handlers

import (
	"fmt"
	"os"
	"trocup-article/services"

	"github.com/gofiber/fiber/v2"
)

type StatusUpdateRequest struct {
	ArticleIDs []string `json:"articleIds"`
	Status     string   `json:"status"`
}

func UpdateArticlesStatus(c *fiber.Ctx) error {
	// Check origin
	origin := c.Get("Origin")
	allowedOrigin := os.Getenv("TRANSACTION_NETWORK")
	
	if origin != allowedOrigin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": fmt.Sprintf("Origin not allowed: %s", origin),
		})
	}

	// Parse request body
	var req StatusUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate request
	if len(req.ArticleIDs) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No article IDs provided",
		})
	}

	// Update status for all articles at once
	err := services.UpdateArticlesStatus(req.ArticleIDs, req.Status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "All articles updated successfully",
	})
} 