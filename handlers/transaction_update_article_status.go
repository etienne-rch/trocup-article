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

func TransactionUpdateArticlesStatus(c *fiber.Ctx) error {
	// Fix the origin variables by moving declarations before use
	origin := c.Get("Origin")
	allowedOrigin := os.Getenv("TRANSACTION_NETWORK")
	
	fmt.Printf("Incoming request from Origin: %s\n", origin)
	fmt.Printf("Allowed Origin from env: %s\n", allowedOrigin)
	
	if origin != allowedOrigin {
		fmt.Printf("Origin mismatch - Request denied for: %s\n", origin)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": fmt.Sprintf("Origin not allowed: %s", origin),
		})
	}

	// Parse request body
	var req StatusUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		fmt.Printf("Failed to parse request body: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Log request details
	fmt.Printf("Request details - ArticleIDs: %v, Status: %s\n", req.ArticleIDs, req.Status)

	// Validate request
	if len(req.ArticleIDs) == 0 {
		fmt.Println("Request rejected: No article IDs provided")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No article IDs provided",
		})
	}

	// Update status for all articles at once
	response, err := services.UpdateArticlesStatus(req.ArticleIDs, req.Status)
	if err != nil {
		fmt.Printf("Error updating articles: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fmt.Printf("Successfully updated %d articles to status: %s\n", len(req.ArticleIDs), req.Status)
	return c.Status(fiber.StatusOK).JSON(response)
} 