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

	// Nettoyer la base de données avant de commencer le test pour éviter tout conflit avec les anciennes données
	config.CleanUpTestDatabase("test_db")

	// Définir la route pour récupérer les articles avec le handler correspondant
	app.Get("/articles", handlers.GetArticles)

	// Initialisation des dates pour les champs ManufactureDate et PurchaseDate
	manufactureDate := time.Date(2023, 10, 8, 0, 0, 0, 0, time.UTC)
	purchaseDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now()

	// Déclarer des pointeurs pour Brand et Model pour correspondre au modèle
	brand1 := "Test Brand"
	model1 := "Test Model"
	brand2 := "Another Brand"
	model2 := "Another Model"

	// Créer deux articles de test avec des valeurs de date et de texte pour les différents champs
	articles := []models.Article{
		{
			ID:              primitive.NewObjectID(),
			Version:         1,
			Owner:           "user_2myWlPeCdykAojnWNwkzUqV3lp9", // ID utilisateur simulé
			AdTitle:         "Test Article 1",
			Brand:           &brand1, // Pointeur vers Brand
			Model:           &model1, // Pointeur vers Model
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
			// Créer un pointeur pour Dimensions
			Dimensions: &models.Dimensions{
				Length: 10,
				Width:  5,
				Height: 1,
				Weight: 0.5,
			},
		},
		{
			ID:              primitive.NewObjectID(),
			Version:         1,
			Owner:           "user_2myWlPeCdykAojnWNwkzUqV3lp8", // ID d'un autre utilisateur simulé
			AdTitle:         "Test Article 2",
			Brand:           &brand2, // Pointeur vers un autre Brand
			Model:           &model2, // Pointeur vers un autre Model
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
			// Créer un pointeur pour Dimensions
			Dimensions: &models.Dimensions{
				Length: 30,
				Width:  20,
				Height: 2,
				Weight: 0.3,
			},
		},
	}

	// Insérer les articles créés dans la base de données de test
	for _, article := range articles {
		config.ArticleCollection.InsertOne(context.TODO(), article)
	}

	// Créer une requête GET pour récupérer les articles
	req := httptest.NewRequest("GET", "/articles", nil)
	resp, _ := app.Test(req, -1)

	// Vérifier que la réponse HTTP renvoie un statut 200 OK
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Décoder le JSON de la réponse pour obtenir les articles et les métadonnées
	var response struct {
		Skip     int              `json:"skip"`
		Limit    int              `json:"limit"`
		Articles []models.Article `json:"articles"`
	}

	err := json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err, "Failed to decode response body")

	// Vérifier que le nombre d'articles retournés correspond à ceux que j'ai insérés dans la base de données
	assert.Equal(t, len(articles), len(response.Articles), "Number of returned articles does not match expected count")

	// Comparer les champs des articles retournés avec ceux que j'ai insérés
	for i, article := range articles {
		assert.Equal(t, article.AdTitle, response.Articles[i].AdTitle)

		// Vérification des champs Brand et Model qui sont des pointeurs
		assert.NotNil(t, response.Articles[i].Brand)
		assert.Equal(t, *article.Brand, *response.Articles[i].Brand) // Comparer les valeurs pointées
		assert.NotNil(t, response.Articles[i].Model)
		assert.Equal(t, *article.Model, *response.Articles[i].Model)

		// Vérifier d'autres champs comme le prix, l'état, le statut, etc.
		assert.Equal(t, article.Price, response.Articles[i].Price)
		assert.Equal(t, article.State, response.Articles[i].State)
		assert.Equal(t, article.Status, response.Articles[i].Status)
		assert.Equal(t, article.Category, response.Articles[i].Category)

		// Comparer les dates pour ManufactureDate et PurchaseDate
		assert.True(t, article.ManufactureDate.Equal(response.Articles[i].ManufactureDate), "expected ManufactureDate to be %v, got %v", article.ManufactureDate, response.Articles[i].ManufactureDate)
		assert.True(t, article.PurchaseDate.Equal(response.Articles[i].PurchaseDate), "expected PurchaseDate to be %v, got %v", article.PurchaseDate)
	}

	// Nettoyage de la base de données après chaque test
	defer config.CleanUpTestDatabase("test_db")
}
