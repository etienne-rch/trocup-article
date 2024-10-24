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

func TestGetArticleByID(t *testing.T) {
	app := fiber.New()

	// Nettoyer la base de données avant le test
	config.CleanUpTestDatabase("test_db")

	// Utiliser le handler pour récupérer un article par ID
	app.Get("/articles/:id", handlers.GetArticleByID)

	// Initialiser des dates sous forme de time.Time
	manufactureDate := time.Date(2023, 10, 8, 0, 0, 0, 0, time.UTC)
	purchaseDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now()

	// Créer des pointeurs pour Brand et Model
	brand := "Test Brand"
	model := "Test Model"

	// Créer un article de test
	article := models.Article{
		ID:              primitive.NewObjectID(),
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

	// Créer une requête GET pour récupérer l'article par ID
	req := httptest.NewRequest("GET", "/articles/"+article.ID.Hex(), nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)

	// Vérifier que le statut de la réponse est 200 OK
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Décoder la réponse JSON pour vérifier le contenu de l'article
	var returnedArticle models.Article
	err = json.NewDecoder(resp.Body).Decode(&returnedArticle)
	assert.NoError(t, err)

	// Vérifications
	assert.Equal(t, article.ID, returnedArticle.ID)
	assert.Equal(t, article.AdTitle, returnedArticle.AdTitle)

	// Vérifier les champs Brand et Model (qui sont des pointeurs)
	assert.NotNil(t, returnedArticle.Brand)
	assert.Equal(t, *article.Brand, *returnedArticle.Brand)
	assert.NotNil(t, returnedArticle.Model)
	assert.Equal(t, *article.Model, *returnedArticle.Model)

	// Vérifier d'autres champs comme le prix, l'état, le statut, etc.
	assert.Equal(t, article.Price, returnedArticle.Price)
	assert.Equal(t, article.State, returnedArticle.State)
	assert.Equal(t, article.Status, returnedArticle.Status)
	assert.Equal(t, article.Category, returnedArticle.Category)

	// Comparaison des dates sous forme de time.Time
	assert.True(t, article.ManufactureDate.Equal(returnedArticle.ManufactureDate), "expected ManufactureDate to be %v, got %v", article.ManufactureDate, returnedArticle.ManufactureDate)
	assert.True(t, article.PurchaseDate.Equal(returnedArticle.PurchaseDate), "expected PurchaseDate to be %v, got %v", article.PurchaseDate)

	// Nettoyage de la base de données après chaque test
	defer config.CleanUpTestDatabase("test_db")
}
