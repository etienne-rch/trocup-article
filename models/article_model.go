package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dimensions struct {
	Length float64 `json:"length" bson:"length"`
	Width  float64 `json:"width" bson:"width"`
	Height float64 `json:"height" bson:"height"`
	Weight float64 `json:"weight" bson:"weight"`
}

type Article struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Version         int                `json:"version" bson:"version"`
	Owner           primitive.ObjectID `json:"owner" bson:"owner"`
	AdTitle         string             `json:"adTitle" bson:"adTitle"`
	Brand           string             `json:"brand" bson:"brand"`
	Model           string             `json:"model" bson:"model"`
	Description     string             `json:"description" bson:"description"`
	Price           int                `json:"price" bson:"price"`
	ManufactureDate time.Time          `json:"manufactureDate" bson:"manufactureDate"`
	PurchaseDate    time.Time          `json:"purchaseDate" bson:"purchaseDate"`
	State           string             `json:"state" bson:"state"`
	Status          string             `json:"status" bson:"status"`
	ImageUrls       []string           `json:"imageUrls" bson:"imageUrls"`
	CreatedAt       time.Time          `json:"createdAt" bson:"createdAt"`
	LastModified    time.Time          `json:"lastModified" bson:"lastModified"`
	Category        string             `json:"category" bson:"category"`
	SubCategory     string             `json:"subCategory" bson:"subCategory"`
	DeliveryType    []string           `json:"deliveryType" bson:"deliveryType"`
	Dimensions      Dimensions         `json:"dimensions" bson:"dimensions"`
}
