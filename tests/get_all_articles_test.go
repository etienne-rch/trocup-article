package tests

import (
	"context"
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

	app.Get("/articles", handlers.GetArticles)

	articles := []models.Article{
		{
			ID:              primitive.NewObjectID(),
			Version:         1,
			Owner:           primitive.NewObjectID(),
			AdTitle:         "Test Article 1",
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
		},
		{
			ID:              primitive.NewObjectID(),
			Version:         1,
			Owner:           primitive.NewObjectID(),
			AdTitle:         "Test Article 2",
			Brand:           "Another Brand",
			Model:           "Another Model",
			Description:     "Another Description",
			Price:           200,
			ManufactureDate: time.Now(),
			PurchaseDate:    time.Now(),
			State:           "Used",
			Status:          "Sold",
			ImageUrls:       []string{"http://example.com/image2.jpg", "http://example.com/image3.jpg"},
			CreatedAt:       time.Now(),
			LastModified:    time.Now(),
			Category:        "Clothing",
			SubCategory:     "Shirts",
			DeliveryType:    []string{"Delivery"},
			Dimensions: models.Dimensions{
				Length: 30,
				Width:  20,
				Height: 2,
				Weight: 0.3,
			},
		},
	}

	for _, article := range articles {
		config.ArticleCollection.InsertOne(context.TODO(), article)
	}

	req := httptest.NewRequest("GET", "/articles", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Cleanup after each test
	defer config.CleanUpTestDatabase("test_db")
}
