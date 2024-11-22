package handlers

import (
	"log"
	"trocup-article/models"
	"trocup-article/repository"
	"trocup-article/services"

	"github.com/gofiber/fiber/v2"
)

func CreateArticle(c *fiber.Ctx) error {
	// Get the JWT token from the request header
	token := c.Get("Authorization")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "No authorization token provided",
		})
	}

	// Get the user ID from the context (set by ClerkAuthMiddleware)
	clerkUserId := c.Locals("clerkUserId").(string)
	if clerkUserId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User ID not found in context",
		})
	}

	// Parse the article from request body
	article := new(models.Article)
	if err := c.BodyParser(article); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse article",
		})
	}

	// Set the owner to the authenticated user's ID
	article.Owner = clerkUserId

	// Save the article to database
	savedArticle, err := repository.CreateArticle(article)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create article",
		})
	}

	log.Printf("Updating user service for clerkUserId: %s", clerkUserId)

	err = services.GetUserService().UpdateUserArticles(
		clerkUserId,
		savedArticle.ID.Hex(),
		article.Price,
		token,
	)

	if err != nil {
		log.Printf("❌ Error updating user service: %v", err)

		// Rollback article creation
		if deleteErr := services.DeleteArticle(savedArticle.ID.Hex()); deleteErr != nil {
			log.Printf("❌ Failed to rollback article creation: %v", deleteErr)
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user information",
		})
	}

	log.Printf("Saved article : %+v", savedArticle)

	return c.Status(fiber.StatusCreated).JSON(savedArticle)
}
