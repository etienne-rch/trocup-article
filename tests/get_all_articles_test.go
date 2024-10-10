package tests

import (
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

func TestGetArticles(t *testing.T) {
	app := fiber.New()

	// Nettoyer la base de données avant le test
	config.CleanUpTestDatabase("test_db")

	// Utiliser le handler pour récupérer les articles
	app.Get("/articles", handlers.GetArticles)

	// Initialiser des dates sous forme de time.Time
	manufactureDate := time.Date(2023, 10, 8, 0, 0, 0, 0, time.UTC)
	purchaseDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now()

	// Définir des articles de test avec des dates sous forme de time.Time
	articles := []models.Article{
		{
			ID:              primitive.NewObjectID(),
			Version:         1,
			Owner:           "user_2myWlPeCdykAojnWNwkzUqV3lp9", // ID utilisateur sous forme de chaîne
			AdTitle:         "Test Article 1",
			Brand:           "Test Brand",
			Model:           "Test Model",
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
			Dimensions: models.Dimensions{
				Length: 10,
				Width:  5,
				Height: 1,
				Weight: 0.5,
			},
		},
		{
			ID:              primitive.NewObjectID(),
			Version:         1,
			Owner:           "user_2myWlPeCdykAojnWNwkzUqV3lp8", // Un autre utilisateur
			AdTitle:         "Test Article 2",
			Brand:           "Another Brand",
			Model:           "Another Model",
			Description:     "Another Description",
			Price:           200,
			ManufactureDate: manufactureDate,
			PurchaseDate:    purchaseDate,
			State:           "USED",
			Status:          "SOLD",
			ImageUrls:       []string{"http://example.com/image2.jpg", "http://example.com/image3.jpg"},
			CreatedAt:       now,
			LastModified:    now,
			Category:        "CLOTHING",
			SubCategory:     "Shirts",
			DeliveryType:    "SHIPPING",
			Dimensions: models.Dimensions{
				Length: 30,
				Width:  20,
				Height: 2,
				Weight: 0.3,
			},
		},
	}

	// Insérer les articles dans la base de données de test
	for _, article := range articles {
		config.ArticleCollection.InsertOne(context.TODO(), article)
	}

	// Créer une requête GET pour récupérer les articles
	req := httptest.NewRequest("GET", "/articles", nil)
	resp, _ := app.Test(req, -1)

	// Vérifier que le statut de la réponse est 200 OK
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Décoder la réponse JSON
	var returnedArticles []models.Article
	err := json.NewDecoder(resp.Body).Decode(&returnedArticles)
	assert.NoError(t, err)

	// Vérifier que le nombre d'articles retournés correspond à ceux insérés
	assert.Equal(t, len(articles), len(returnedArticles))

	// Vérifier que le contenu des articles correspond à ceux insérés
	for i, article := range articles {
		assert.Equal(t, article.AdTitle, returnedArticles[i].AdTitle)
		assert.Equal(t, article.Brand, returnedArticles[i].Brand)
		assert.Equal(t, article.Price, returnedArticles[i].Price)
		assert.Equal(t, article.State, returnedArticles[i].State)
		assert.Equal(t, article.Status, returnedArticles[i].Status)
		assert.Equal(t, article.Category, returnedArticles[i].Category)
		// Comparaison des dates sous forme de time.Time
		assert.True(t, article.ManufactureDate.Equal(returnedArticles[i].ManufactureDate), "expected ManufactureDate to be %v, got %v", article.ManufactureDate, returnedArticles[i].ManufactureDate)
		assert.True(t, article.PurchaseDate.Equal(returnedArticles[i].PurchaseDate), "expected PurchaseDate to be %v, got %v", article.PurchaseDate)
	}

	// Nettoyage de la base de données après chaque test
	defer config.CleanUpTestDatabase("test_db")
}
