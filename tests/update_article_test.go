package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"trocup-article/config"
	"trocup-article/handlers"
	"trocup-article/models"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUpdateArticle(t *testing.T) {
	app := fiber.New()

	// Mock le middleware Clerk pour simuler l'authentification
	app.Use(func(c *fiber.Ctx) error {
		// Simuler un utilisateur authentifié en définissant un faux ID utilisateur dans le contexte
		c.Locals("clerkUserId", "user_2myWlPeCdykAojnWNwkzUqV3lp9") // ID simulé
		return c.Next()
	})

	// Utiliser le handler pour mettre à jour un article
	app.Put("/articles/:id", handlers.UpdateArticle)

	// Initialiser des dates sous forme de time.Time
	manufactureDate := time.Date(2023, 10, 8, 0, 0, 0, 0, time.UTC)
	purchaseDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now()

	// Pointeurs pour Brand et Model
	brand := "Test Brand"
	model := "Test Model"

	// Insérer un article de test
	article := models.Article{
		ID:              primitive.NewObjectID(),
		Version:         1,
		Owner:           "user_2myWlPeCdykAojnWNwkzUqV3lp9", // ID utilisateur sous forme de chaîne
		AdTitle:         "Test Article",
		Brand:           &brand, // Pointeur pour Brand
		Model:           &model, // Pointeur pour Model
		Description:     "Test Description",
		Price:           100,
		ManufactureDate: manufactureDate,
		PurchaseDate:    purchaseDate,
		State:           "NEW",
		Status:          "AVAILABLE",
		ImageUrls:       []string{"http://example.com/image1.jpg"},
		CreatedAt:       now,
		LastModified:    now,
		Category:        "ELECTRONICS",
		SubCategory:     "Smartphones",
		DeliveryType:    "PICKUP",
		// Utiliser un pointeur pour Dimensions
		Dimensions: &models.Dimensions{
			Length: 10,
			Width:  5,
			Height: 1,
			Weight: 0.5,
		},
	}

	// Insérer l'article dans la base de données de test
	_, err := config.ArticleCollection.InsertOne(context.TODO(), article)
	assert.NoError(t, err)

	// Pointeurs pour les nouvelles valeurs de Brand et Model après la mise à jour
	updatedBrand := "Updated Brand"
	updatedModel := "Updated Model"

	// Modifier les données de l'article
	updatedArticle := models.Article{
		AdTitle:      "Updated Article",
		Brand:        &updatedBrand, // Pointeur pour le Brand mis à jour
		Model:        &updatedModel, // Pointeur pour le Model mis à jour
		Description:  "Updated Description",
		Price:        200,
		State:        "USED",
		Status:       "SOLD",
		ImageUrls:    []string{"http://example.com/image2.jpg"},
		Category:     "CLOTHING",
		SubCategory:  "Shirts",
		DeliveryType: "SHIPPING",
		// Pointeur pour les Dimensions mises à jour
		Dimensions: &models.Dimensions{
			Length: 20,
			Width:  10,
			Height: 2,
			Weight: 1,
		},
	}

	// Convertir l'article modifié en JSON
	jsonArticle, _ := json.Marshal(updatedArticle)
	req := httptest.NewRequest("PUT", "/articles/"+article.ID.Hex(), bytes.NewReader(jsonArticle))
	req.Header.Set("Content-Type", "application/json")

	// Envoyer la requête de mise à jour
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)

	// Vérifier que le statut de la réponse est 200 OK
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Décoder la réponse pour vérifier les modifications
	var returnedArticle models.Article
	err = json.NewDecoder(resp.Body).Decode(&returnedArticle)
	assert.NoError(t, err)

	// Vérifications
	assert.Equal(t, updatedArticle.AdTitle, returnedArticle.AdTitle)
	assert.Equal(t, updatedArticle.Price, returnedArticle.Price)
	assert.Equal(t, updatedArticle.State, returnedArticle.State)
	assert.Equal(t, updatedArticle.Status, returnedArticle.Status)
	assert.Equal(t, updatedArticle.Category, returnedArticle.Category)

	// Vérifier que les champs Brand et Model ont bien été mis à jour
	assert.NotNil(t, returnedArticle.Brand)
	assert.Equal(t, *updatedArticle.Brand, *returnedArticle.Brand)
	assert.NotNil(t, returnedArticle.Model)
	assert.Equal(t, *updatedArticle.Model, *returnedArticle.Model)

	// Nettoyage de la base de données après chaque test
	defer config.CleanUpTestDatabase("test_db")
}
