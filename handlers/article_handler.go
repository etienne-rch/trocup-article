package handlers

import (
	"trocup-article/models"
	"trocup-article/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetupRoutes(app *fiber.App) {
    app.Post("/articles", createArticle)
    app.Get("/articles", getArticles)
	app.Get("/articles/:id", getArticleByID)
}

func createArticle(c *fiber.Ctx) error {
    var article models.Article
    if err := c.BodyParser(&article); err != nil {
        return c.Status(400).SendString(err.Error())
    }

    if err := services.CreateArticle(&article); err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.JSON(article)
}

func getArticles(c *fiber.Ctx) error {
    articles, err := services.GetArticles()
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.JSON(articles)
}

func getArticleByID(c *fiber.Ctx) error {
    id, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return c.Status(400).SendString("Invalid ID")
    }

    article, err := services.GetArticleByID(id)
    if err != nil {
        return c.Status(404).SendString("Article not found")
    }
    return c.JSON(article)
}
