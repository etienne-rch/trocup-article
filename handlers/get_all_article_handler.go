package handlers

import (
	"os"
	"strconv"
	"trocup-article/services"

	"github.com/gofiber/fiber/v2"
)

func GetArticles(c *fiber.Ctx) error {
	// Get default skip and limit from environment variables
	defaultSkip, _ := strconv.ParseInt(os.Getenv("DEFAULT_SKIP"), 10, 64)
	defaultLimit, _ := strconv.ParseInt(os.Getenv("DEFAULT_LIMIT"), 10, 64)

	// If env variables aren't set or parsing fails, fallback to hardcoded defaults
	if defaultSkip < 0 {
		defaultSkip = 0
	}
	if defaultLimit <= 0 {
		defaultLimit = 100
	}

	// Get query parameters
	skipParam := c.Query("skip", strconv.FormatInt(defaultSkip, 10))    // Use default from env if not provided
	limitParam := c.Query("limit", strconv.FormatInt(defaultLimit, 10)) // Use default from env if not provided
	category := c.Query("category", "")
	status := c.Query("status", "")

	skip, err := strconv.ParseInt(skipParam, 10, 64)
	if err != nil || skip < 0 {
		skip = defaultSkip
	}

	limit, err := strconv.ParseInt(limitParam, 10, 64)
	if err != nil || limit <= 0 {
		limit = defaultLimit
	}

	articles, hasNext, err := services.GetAllArticles(skip, limit, category, status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"skip":     skip,
		"limit":    limit,
		"hasNext":  hasNext,
		"articles": articles,
	})
}
