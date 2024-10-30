package tests

import (
	"bytes"
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
)

func TestCreateArticle(t *testing.T) {
	app := fiber.New()

	// Mock le middleware ClerkAuthMiddleware pour ignorer l'authentification
	app.Use(func(c *fiber.Ctx) error {
		// Simuler un utilisateur authentifié en définissant un faux ID utilisateur dans le contexte
		c.Locals("clerkUserId", "user_2myWlPeCdykAojnWNwkzUqV3lp9")
		return c.Next()
	})

	// Utiliser le handler pour créer un article
	app.Post("/articles", handlers.CreateArticle)

	// Initialiser des dates sous forme de time.Time pour manufactureDate et purchaseDate
	manufactureDate := time.Date(2023, 10, 8, 0, 0, 0, 0, time.UTC)
	purchaseDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now()

	// Définition des dimensions
	dimensions := models.Dimensions{
		Length: 10.0,
		Width:  5.0,
		Height: 1.0,
		Weight: 0.5,
	}

	// Créer des valeurs pour les champs `Brand` et `Model` qui sont des pointeurs
	brand := "Test Brand"
	model := "Test Model"

	// Définir un article de test avec des dates sous forme de time.Time
	article := models.Article{
		Owner:           "user_2myWlPeCdykAojnWNwkzUqV3lp9", // ID utilisateur en string (Clerk ID simulé)
		AdTitle:         "Test Article",
		Brand:           &brand,
		Model:           &model,
		Description:     "Test Description",
		Price:           100,
		ManufactureDate: manufactureDate,
		PurchaseDate:    purchaseDate,
		State:           "NEW",
		Status:          "AVAILABLE",
		ImageUrls:       []string{"http://example.com/image1.jpg", "http://example.com/image2.jpg"},
		CreatedAt:       now,
		LastModified:    now,
		Category:        "ELECTRONICS",
		SubCategory:     "Smartphones",
		DeliveryType:    "PICKUP",
		Dimensions:      &dimensions,
		Address: models.Address{
			City:     "Paris",
			Postcode: "75001",
			Citycode: "12345",
			GeoPoints: models.GeoPoints{
				Type:        "Point",
				Coordinates: []float64{2.3522, 48.8566}, // Coordonnées pour Paris
			},
		},
	}

	// Convertir l'article en JSON
	jsonArticle, _ := json.Marshal(article)

	// Créer une requête POST pour créer l'article
	req := httptest.NewRequest("POST", "/articles", bytes.NewReader(jsonArticle))
	req.Header.Set("Content-Type", "application/json")

	// Tester la requête
	resp, _ := app.Test(req, -1)

	// Vérifier que le statut de la réponse est 201 Created
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Décoder la réponse pour vérifier l'article créé
	var createdArticle models.Article
	json.NewDecoder(resp.Body).Decode(&createdArticle)

	// Vérifications
	assert.Equal(t, article.AdTitle, createdArticle.AdTitle)

	// Vérifier les champs optionnels (pointeurs)
	assert.NotNil(t, createdArticle.Brand)
	assert.Equal(t, *article.Brand, *createdArticle.Brand)
	assert.NotNil(t, createdArticle.Model)
	assert.Equal(t, *article.Model, *createdArticle.Model)

	// Vérifier les autres champs
	assert.Equal(t, article.Price, createdArticle.Price)
	assert.Equal(t, article.State, createdArticle.State)
	assert.Equal(t, article.Status, createdArticle.Status)
	assert.Equal(t, article.Category, createdArticle.Category)
	assert.Equal(t, article.SubCategory, createdArticle.SubCategory)

	// Comparaison des dates sous forme de time.Time
	assert.True(t, article.ManufactureDate.Equal(createdArticle.ManufactureDate), "expected ManufactureDate to be %v, got %v", article.ManufactureDate, createdArticle.ManufactureDate)
	assert.True(t, article.PurchaseDate.Equal(createdArticle.PurchaseDate), "expected PurchaseDate to be %v, got %v", article.PurchaseDate, createdArticle.PurchaseDate)

	// Nettoyage de la base de données après le test
	defer config.CleanUpTestDatabase("test_db")
}
