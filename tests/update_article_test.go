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

	app.Put("/articles/:id", handlers.UpdateArticle)

	// Insert an article for testing
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

	// Modify the article data
	updatedArticle := models.Article{
		AdTitle:      "Updated Article",
		Brand:        "Updated Brand",
		Model:        "Updated Model",
		Description:  "Updated Description",
		Price:        200,
		State:        "Used",
		Status:       "Sold",
		ImageUrls:    []string{"http://example.com/image2.jpg"},
		Category:     "Clothing",
		SubCategory:  "Shirts",
		DeliveryType: []string{"Delivery"},
		Dimensions: models.Dimensions{
			Length: 20,
			Width:  10,
			Height: 2,
			Weight: 1,
		},
	}

	jsonArticle, _ := json.Marshal(updatedArticle)
	req := httptest.NewRequest("PUT", "/articles/"+article.ID.Hex(), bytes.NewReader(jsonArticle))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var returnedArticle models.Article
	json.NewDecoder(resp.Body).Decode(&returnedArticle)

	assert.Equal(t, updatedArticle.AdTitle, returnedArticle.AdTitle)
	assert.Equal(t, updatedArticle.Price, returnedArticle.Price)
	assert.Equal(t, updatedArticle.State, returnedArticle.State)
	assert.Equal(t, updatedArticle.Status, returnedArticle.Status)

	// Cleanup after each test
	defer config.CleanUpTestDatabase("test_db")
}
