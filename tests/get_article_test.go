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

	app.Get("/articles/:id", handlers.GetArticleByID)

	article := models.Article{
		ID:              primitive.NewObjectID(),
		Version:         1,
		Owner:           primitive.NewObjectID(),
		AdTitle:         "Test Article",
		Brand:           "Test Brand",
		Model:           "Test Model",
		Description:     "Test Description",
		Price:           100,
		ManufactureDate: time.Now(),
		PurchaseDate:    time.Now(),
		State:           "New",
		Status:          "Available",
		ImageUrls:       []string{"http://example.com/image1.jpg"},
		CreatedAt:       time.Now(),
		LastModified:    time.Now(),
		Category:        "Electronics",
		SubCategory:     "Smartphones",
		DeliveryType:    []string{"Pickup", "Delivery"},
		Dimensions: models.Dimensions{
			Length: 10,
			Width:  5,
			Height: 1,
			Weight: 0.5,
		},
	}
	config.ArticleCollection.InsertOne(context.TODO(), article)

	req := httptest.NewRequest("GET", "/articles/"+article.ID.Hex(), nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var returnedArticle models.Article
	json.NewDecoder(resp.Body).Decode(&returnedArticle)

	assert.Equal(t, article.ID, returnedArticle.ID)
	assert.Equal(t, article.AdTitle, returnedArticle.AdTitle)

	// Cleanup after each test
	defer config.CleanUpTestDatabase("test_db")
}
