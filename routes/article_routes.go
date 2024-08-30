package routes

import (
	"fmt"
	"trocup-article/handlers"
	"trocup-article/middleware"

	"github.com/gofiber/fiber/v2"
)

func ArticleRoutes(app *fiber.App) {
	// PUBLIC
	app.Get("/health", handlers.HealthCheck)

	// PRIVATE
	api := app.Group("/api", middleware.AuthMiddleware())

	api.Get("/articles", handlers.GetArticles)
	api.Get("/articles/:id", handlers.GetArticleByID)
	api.Post("/articles", handlers.CreateArticle)
	api.Put("/articles/:id", handlers.UpdateArticle)
	api.Delete("/articles/:id", handlers.DeleteArticle)

	// Add a catch-all route for debugging
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString(fmt.Sprintf("Route not found: %s", c.Path()))
	})
}
