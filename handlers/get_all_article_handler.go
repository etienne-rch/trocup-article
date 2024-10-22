package handlers

import (
	"strconv"
	"trocup-article/services"

	"github.com/gofiber/fiber/v2"
)

func GetArticles(c *fiber.Ctx) error {
	// Retrieve pagination parameters
	skipParam := c.Query("skip", "0")    // Default value: skip = 0
	limitParam := c.Query("limit", "10") // Default value: limit = 10

	// Parse skip and limit
	skip, err := strconv.ParseInt(skipParam, 10, 64)
	if err != nil || skip < 0 {
		skip = 0
	}

	limit, err := strconv.ParseInt(limitParam, 10, 64)
	if err != nil || limit <= 0 {
		limit = 10
	}

	// Extract geo parameters from the query and convert them to float64
	latitudeParam := c.Query("latitude")
	longitudeParam := c.Query("longitude")
	radiusParam := c.Query("radius", "5") // Default radius in km

	latitude, _ := strconv.ParseFloat(latitudeParam, 64)
	longitude, _ := strconv.ParseFloat(longitudeParam, 64)
	radius, _ := strconv.ParseFloat(radiusParam, 64)

	// Call the service to get articles
	articles, hasNext, err := services.GetAllArticles(skip, limit, latitude, longitude, radius)
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
