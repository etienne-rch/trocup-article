package routes

import (
	"trocup-article/handlers"

	"github.com/gofiber/fiber/v2"
)

func ArticleRoutes(app *fiber.App) {
	app.Get("/health", handlers.HealthCheck)

	app.Get("/articles", handlers.GetArticles)
	app.Get("/articles/:id", handlers.GetArticleByID)

	app.Post("/articles", handlers.CreateArticle)

	app.Put("/articles/:id", handlers.UpdateArticle)

	app.Delete("/articles/:id", handlers.DeleteArticle)

}
