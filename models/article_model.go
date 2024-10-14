package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dimensions struct {
	Length float64 `json:"length,omitempty" bson:"length,omitempty"`
	Width  float64 `json:"width,omitempty" bson:"width,omitempty"`
	Height float64 `json:"height,omitempty" bson:"height,omitempty"`
	Weight float64 `json:"weight,omitempty" bson:"weight,omitempty"`
}

type GeoPoints struct {
	Type        string    `json:"type,omitempty" bson:"type,omitempty" validate:"required,eq=Point"`
	Coordinates []float64 `json:"coordinates,omitempty" bson:"coordinates,omitempty" validate:"required"`
}

type Address struct {
	City      string    `json:"city" bson:"city" validate:"required"`
	Postcode  string    `json:"postcode" bson:"postcode" validate:"required"`
	Citycode  string    `json:"citycode" bson:"citycode" validate:"required"`
	GeoPoints GeoPoints `json:"geopoints" bson:"geopoints" validate:"required"`
}

type Article struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Version         int                `json:"version" bson:"version" validate:"required,gt=0"`
	Owner           string             `json:"owner,omitempty" bson:"owner,omitempty"`
	AdTitle         string             `json:"adTitle" bson:"adTitle" validate:"required"`
	Brand           *string            `json:"brand,omitempty" bson:"brand,omitempty"`
	Model           *string            `json:"model,omitempty" bson:"model,omitempty"`
	Description     string             `json:"description" bson:"description" validate:"required"`
	Price           int                `json:"price" bson:"price" validate:"required,gt=0"`
	ManufactureDate time.Time          `json:"manufactureDate" bson:"manufactureDate"`
	PurchaseDate    time.Time          `json:"purchaseDate" bson:"purchaseDate"`
	State           string             `json:"state" bson:"state" validate:"required,oneof=NEW LIKE_NEW VERY_GOOD_CONDITION GOOD_CONDITION FAIR_CONDITION TO_REPAIR"`
	Status          string             `json:"status" bson:"status" validate:"required,oneof=AVAILABLE UNAVAILABLE RESERVED"`
	ImageUrls       []string           `json:"imageUrls" bson:"imageUrls" validate:"required,dive,url"`
	CreatedAt       time.Time          `json:"createdAt" bson:"createdAt"`
	LastModified    time.Time          `json:"lastModified" bson:"lastModified"`
	Category        string             `json:"category" bson:"category" validate:"required"`
	SubCategory     string             `json:"subCategory" bson:"subCategory"`
	DeliveryType    string             `json:"deliveryType" bson:"deliveryType" validate:"required,oneof=PICKUP SHIPPING BOTH"`
	Dimensions      *Dimensions        `json:"dimensions,omitempty" bson:"dimensions,omitempty"`
	Address         Address            `json:"address" bson:"address" validate:"required"`
	Size            *string            `json:"size,omitempty" bson:"size,omitempty"`
}
