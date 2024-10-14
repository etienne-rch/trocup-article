package handlers

import (
	"trocup-article/services"

	"github.com/gofiber/fiber/v2"
)

func DeleteArticle(c *fiber.Ctx) error {
	// Récupérer l'ID de l'article à partir des paramètres d'URL
	id := c.Params("id")

	// Récupérer l'ID utilisateur connecté à partir du contexte (défini par le middleware Clerk)
	clerkUserId := c.Locals("clerkUserId").(string)

	// Récupérer l'article à partir de la base de données via le service GetArticleByID
	article, err := services.GetArticleByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Article not found"})
	}

	// Vérifier que l'utilisateur connecté est bien le propriétaire de l'article
	if article.Owner != clerkUserId {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You do not have permission to delete this article"})
	}

	// Supprimer l'article via le service
	err = services.DeleteArticle(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Retourner un statut 204 No Content si la suppression est réussie
	return c.SendStatus(fiber.StatusNoContent)
}
