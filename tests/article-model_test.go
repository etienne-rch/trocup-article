package tests

import (
	"testing"
	"time"
	"trocup-article/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestArticle(t *testing.T) {
	dimensions := models.Dimensions{
		Length: float64(10.0),
		Width:  float64(5.0),
		Height: float64(1.0),
		Weight: float64(0.5),
	}

	// ID utilisateur Clerk (string)
	ownerID := "user_2myWlPeCdykAojnWNwkzUqV3lp9"

	// Créer des dates sous forme de time.Time
	manufactureDate := time.Date(2023, 10, 9, 0, 0, 0, 0, time.UTC)
	purchaseDate := time.Date(2023, 12, 15, 0, 0, 0, 0, time.UTC)

	article := models.Article{
		ID:              primitive.NewObjectID(), // ID généré
		Version:         1,
		Owner:           ownerID, // Propriétaire est une string
		AdTitle:         "First Article",
		Brand:           "BrandName",
		Model:           "ModelName",
		Description:     "This is the body of the first article.",
		Price:           100,
		ManufactureDate: manufactureDate,
		PurchaseDate:    purchaseDate,
		State:           "NEW",
		Status:          "AVAILABLE",
		ImageUrls:       []string{"url1", "url2"},
		CreatedAt:       time.Now(),
		LastModified:    time.Now(),
		Category:        "ELECTRONICS",
		SubCategory:     "phone",
		DeliveryType:    "PICKUP",
		Dimensions:      dimensions,
	}

	// Test des valeurs
	if article.Owner != ownerID {
		t.Errorf("expected Owner to be %v, got %v", ownerID, article.Owner)
	}
	if article.AdTitle != "First Article" {
		t.Errorf("expected AdTitle to be 'First Article', got %s", article.AdTitle)
	}
	if article.Brand != "BrandName" {
		t.Errorf("expected Brand to be 'BrandName', got %s", article.Brand)
	}
	if article.Model != "ModelName" {
		t.Errorf("expected Model to be 'ModelName', got %s", article.Model)
	}
	if article.Description != "This is the body of the first article." {
		t.Errorf("expected Description to be 'This is the body of the first article.', got %s", article.Description)
	}
	if article.Price != 100 {
		t.Errorf("expected Price to be 100, got %d", article.Price)
	}
	// Comparaison des dates sous forme de time.Time
	if !article.ManufactureDate.Equal(manufactureDate) {
		t.Errorf("expected ManufactureDate to be %v, got %v", manufactureDate, article.ManufactureDate)
	}
	if !article.PurchaseDate.Equal(purchaseDate) {
		t.Errorf("expected PurchaseDate to be %v, got %v", purchaseDate, article.PurchaseDate)
	}
	if article.State != "NEW" {
		t.Errorf("expected State to be 'NEW', got %s", article.State)
	}
	if article.Status != "AVAILABLE" {
		t.Errorf("expected Status to be 'AVAILABLE', got %s", article.Status)
	}
	if len(article.ImageUrls) != 2 || article.ImageUrls[0] != "url1" || article.ImageUrls[1] != "url2" {
		t.Errorf("expected ImageUrls to be ['url1', 'url2'], got %v", article.ImageUrls)
	}
	if article.Category != "ELECTRONICS" {
		t.Errorf("expected Category to be 'ELECTRONICS', got %s", article.Category)
	}
	if article.SubCategory != "phone" {
		t.Errorf("expected SubCategory to be 'phone', got %s", article.SubCategory)
	}
	if article.DeliveryType != "PICKUP" {
		t.Errorf("expected DeliveryType to be 'PICKUP', got %v", article.DeliveryType)
	}
	if article.Dimensions != dimensions {
		t.Errorf("expected Dimensions to be %v, got %v", dimensions, article.Dimensions)
	}
}
