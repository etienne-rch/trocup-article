package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"trocup-article/config"
	"trocup-article/handlers"
	"trocup-article/models"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestDeleteArticle(t *testing.T) {
	app := fiber.New()

	// Nettoyer la base de données avant le test
	config.CleanUpTestDatabase("test_db")

	// Mock le middleware Clerk pour simuler l'authentification
	app.Use(func(c *fiber.Ctx) error {
		// Simuler un utilisateur authentifié en définissant un faux ID utilisateur dans le contexte
		c.Locals("clerkUserId", "user_2myWlPeCdykAojnWNwkzUqV3lp9") // ID simulé
		return c.Next()
	})

	// Utiliser le handler pour supprimer un article
	app.Delete("/articles/:id", handlers.DeleteArticle)

	// Définir des pointeurs pour `Brand` et `Model`
	brand := "Test Brand"
	model := "Test Model"

	// Créer un article de test
	article := models.Article{
		ID:          primitive.NewObjectID(),
		Version:     1,
		Owner:       "user_2myWlPeCdykAojnWNwkzUqV3lp9", // ID utilisateur simulé
		AdTitle:     "Test Article",
		Brand:       &brand,
		Model:       &model,
		Description: "Test Description",
		Price:       100,
		State:       "NEW",
		Status:      "AVAILABLE",
		ImageUrls:   []string{"http://example.com/image1.jpg"},
	}

	// Insérer l'article dans la base de données de test
	result, err := config.ArticleCollection.InsertOne(context.TODO(), article)
	assert.NoError(t, err)

	// Log l'ID inséré pour vérifier
	t.Log("Inserted article with ID:", result.InsertedID.(primitive.ObjectID).Hex())

	// Créer une requête DELETE pour supprimer l'article
	req := httptest.NewRequest("DELETE", "/articles/"+article.ID.Hex(), nil)
	req.Header.Set("Content-Type", "application/json")

	// Envoyer la requête DELETE
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)

	// Vérifier que le statut de la réponse est 204 No Content
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	// Vérifier que l'article a bien été supprimé
	var deletedArticle models.Article
	err = config.ArticleCollection.FindOne(context.TODO(), bson.M{"_id": article.ID}).Decode(&deletedArticle)

	// Vérifier que l'erreur est bien une erreur indiquant que le document n'a pas été trouvé
	assert.Equal(t, mongo.ErrNoDocuments, err, "expected ErrNoDocuments, got %v", err)

	// Nettoyage de la base de données après chaque test
	defer config.CleanUpTestDatabase("test_db")
}
