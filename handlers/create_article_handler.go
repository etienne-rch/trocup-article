package handlers

import (
	"log"
	"trocup-article/models"
	"trocup-article/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateArticle(c *fiber.Ctx) error {
	var validate = validator.New()
	article := new(models.Article)

	// Parse le corps de la requête JSON dans le modèle Article
	if err := c.BodyParser(&article); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Récupérer l'ID utilisateur connecté à partir du contexte (défini par le middleware Clerk)
	clerkUserId := c.Locals("clerkUserId").(string)
	article.Owner = clerkUserId

	// Validation des données reçues via le validateur
	if err := validate.Struct(article); err != nil {
		log.Printf("Validation error: %v", err)
		// Retourner une erreur si la validation échoue
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Créer l'article
	createdArticle, err := services.CreateArticle(article)
	if err != nil {
		log.Printf("Error creating article: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Retourner l'article créé avec un statut 201 (Created)
	return c.Status(fiber.StatusCreated).JSON(createdArticle)
}
