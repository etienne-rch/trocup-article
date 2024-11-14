package routes

import (
	"fmt"
	"trocup-article/handlers"
	"trocup-article/middleware"

	"github.com/gofiber/fiber/v2"
)

func ArticleRoutes(app *fiber.App) {
	// Routes publiques : accessibles sans authentification
	public := app.Group("/api")

	public.Get("/health", handlers.HealthCheck)
	public.Get("/articles", handlers.GetArticles)
	public.Get("/articles/:id", handlers.GetArticleByID)


	// Routes protégées : accessibles uniquement avec authentification
	protected := app.Group("/api/protected", middleware.ClerkAuthMiddleware)

	protected.Post("/articles", handlers.CreateArticle)
	protected.Put("/articles/:id", handlers.UpdateArticle)
	protected.Delete("/articles/:id", handlers.DeleteArticle)
	protected.Patch("/articles/status", handlers.UpdateArticlesStatus)

	// Add a catch-all route for debugging
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString(fmt.Sprintf("Route not found: %s", c.Path()))
	})
}
