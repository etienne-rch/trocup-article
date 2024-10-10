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

	// Routes publiques : accessibles sans authentification
	app.Get("/articles", handlers.GetArticles)
	app.Get("/articles/:id", handlers.GetArticleByID)

	// PRIVATE : Routes protégées par le middleware ClerkAuthMiddleware
	api := app.Group("/api", middleware.ClerkAuthMiddleware)

	api.Post("/articles", handlers.CreateArticle)
	api.Put("/articles/:id", handlers.UpdateArticle)
	api.Delete("/articles/:id", handlers.DeleteArticle)

	// Add a catch-all route for debugging
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString(fmt.Sprintf("Route not found: %s", c.Path()))
	})
}
