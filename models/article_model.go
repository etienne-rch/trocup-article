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
	Label     string    `json:"label,omitempty" bson:"label,omitempty"`
	Street     string    `json:"street,omitempty" bson:"street,omitempty"`
	City       string    `json:"city,omitempty" bson:"city,omitempty"`
	Postcode   string       `json:"postcode,omitempty" bson:"postcode,omitempty"`
	Citycode   string       `json:"citycode,omitempty" bson:"citycode,omitempty"`
	Floor      int       `json:"floor,omitempty" bson:"floor,omitempty"`
	Extra      string    `json:"extra,omitempty" bson:"extra,omitempty"`
	GeoPoints  GeoPoints `json:"geopoints,omitempty" bson:"geopoints,omitempty"`
}

type Article struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Owner           string             `json:"owner" bson:"owner" validate:"required"`
	AdTitle         string             `json:"adTitle" bson:"adTitle" validate:"required"`
	Brand           *string            `json:"brand,omitempty" bson:"brand,omitempty"`
	Model           *string            `json:"model,omitempty" bson:"model,omitempty"`
	Description     string             `json:"description" bson:"description" validate:"required"`
	Price           float64           `json:"price" bson:"price" validate:"required,gt=0"`
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
