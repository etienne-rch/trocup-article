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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateArticle(t *testing.T) {
	app := fiber.New()

	app.Post("/articles", handlers.CreateArticle)

	article := models.Article{
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
		ImageUrls:       []string{"http://example.com/image1.jpg", "http://example.com/image2.jpg"},
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

	jsonArticle, _ := json.Marshal(article)
	req := httptest.NewRequest("POST", "/articles", bytes.NewReader(jsonArticle))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var createdArticle models.Article
	json.NewDecoder(resp.Body).Decode(&createdArticle)

	assert.Equal(t, article.AdTitle, createdArticle.AdTitle)
	assert.Equal(t, article.Brand, createdArticle.Brand)
	assert.Equal(t, article.Price, createdArticle.Price)
	assert.Equal(t, article.State, createdArticle.State)
	assert.Equal(t, article.Status, createdArticle.Status)
	assert.Equal(t, article.Category, createdArticle.Category)
	assert.Equal(t, article.SubCategory, createdArticle.SubCategory)

	// Cleanup after each test
	defer config.CleanUpTestDatabase("test_db")
}
