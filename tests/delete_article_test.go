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

func TestDeleteArticle(t *testing.T) {
	app := fiber.New()

	app.Delete("/articles/:id", handlers.DeleteArticle)

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

	req := httptest.NewRequest("DELETE", "/articles/"+article.ID.Hex(), nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	// Verify that the article has been deleted
	var deletedArticle models.Article
	err := config.ArticleCollection.FindOne(context.TODO(), primitive.M{"_id": article.ID}).Decode(&deletedArticle)
	assert.Error(t, err) // Expect an error because the article should not be found

	// Cleanup after each test
	defer config.CleanUpTestDatabase("test_db")
}
