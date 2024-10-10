package services

import (
	"fmt"
	"time"
	"trocup-article/models"
	"trocup-article/repository"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdateArticle passe la mise à jour de l'article à la couche repository
func UpdateArticle(c *fiber.Ctx, articleID string) (*models.Article, error) {
	// Convertir articleID en ObjectID MongoDB
	objID, err := primitive.ObjectIDFromHex(articleID)
	if err != nil {
		return nil, fmt.Errorf("invalid article ID")
	}

	// Récupérer les nouvelles valeurs depuis le body
	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return nil, fmt.Errorf("invalid request body")
	}

	// Mettre à jour la date de modification
	updates["lastModified"] = time.Now()

	// Appel au repository pour effectuer la mise à jour partielle
	updatedArticle, err := repository.UpdateArticle(objID, updates) // Utiliser objID ici
	if err != nil {
		return nil, err
	}

	return updatedArticle, nil
}
