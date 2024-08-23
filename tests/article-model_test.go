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

	id, err := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")
	if err != nil {
		t.Fatalf("failed to create ObjectID: %v", err)
	}

	now := time.Now()

	article := models.Article{
		ID:          id,
		Version:     1,
		Owner:       id,
		AdTitle:     "First Article",
		Brand:       "BrandName",
		Model:       "ModelName",
		Description: "This is the body of the first article.",
		Price:       100,
		// ManufactureDate: now,
		// PurchaseDate:    now,
		State:        "new",
		Status:       "available",
		ImageUrls:    []string{"url1", "url2"},
		CreatedAt:    now,
		LastModified: now,
		Category:     "electronics",
		SubCategory:  "phone",
		DeliveryType: []string{"standard"},
		Dimensions:   dimensions,
	}

	if article.ID != id {
		t.Errorf("expected ID to be %v, got %v", id, article.ID)
	}
	if article.Version != 1 {
		t.Errorf("expected Version to be 1, got %d", article.Version)
	}
	if article.Owner != id {
		t.Errorf("expected Owner to be %v, got %v", id, article.Owner)
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
	// if article.ManufactureDate != now {
	// 	t.Errorf("expected ManufactureDate to be %v, got %v", now, article.ManufactureDate)
	// }
	// if article.PurchaseDate != now {
	// 	t.Errorf("expected PurchaseDate to be %v, got %v", now, article.PurchaseDate)
	// }
	if article.State != "new" {
		t.Errorf("expected State to be 'new', got %s", article.State)
	}
	if article.Status != "available" {
		t.Errorf("expected Status to be 'available', got %s", article.Status)
	}
	if len(article.ImageUrls) != 2 || article.ImageUrls[0] != "url1" || article.ImageUrls[1] != "url2" {
		t.Errorf("expected ImageUrls to be ['url1', 'url2'], got %v", article.ImageUrls)
	}
	if article.CreatedAt != now {
		t.Errorf("expected CreatedAt to be %v, got %v", now, article.CreatedAt)
	}
	if article.LastModified != now {
		t.Errorf("expected LastModified to be %v, got %v", now, article.LastModified)
	}
	if article.Category != "electronics" {
		t.Errorf("expected Category to be 'electronics', got %s", article.Category)
	}
	if article.SubCategory != "phone" {
		t.Errorf("expected SubCategory to be 'phone', got %s", article.SubCategory)
	}
	if len(article.DeliveryType) != 1 || article.DeliveryType[0] != "standard" {
		t.Errorf("expected DeliveryType to be ['standard'], got %v", article.DeliveryType)
	}
	if article.Dimensions != dimensions {
		t.Errorf("expected Dimensions to be %v, got %v", dimensions, article.Dimensions)
	}
}
