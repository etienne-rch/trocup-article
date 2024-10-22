package repository

import (
	"context"
	"fmt"
	"trocup-article/config"
	"trocup-article/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAllArticles retrieves articles from the database with optional geo-location filtering
func GetAllArticles(skip, limit int64, latitude, longitude float64, radiusInKm float64) ([]models.Article, bool, error) {
	var articles []models.Article

	// Create find options for pagination
	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)

	// Create base filter
	filter := bson.M{}

	// If latitude and longitude are provided, apply geo-location filter
	if latitude != 0 && longitude != 0 {
		radiusInRadians := radiusInKm / 6378.1 // Convert radius to radians
		filter["location"] = bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": bson.A{
					bson.A{longitude, latitude}, // Coordinates in [longitude, latitude]
					radiusInRadians,             // Radius in radians
				},
			},
		}
	}

	// Count total number of articles based on the filter (with or without geo-location)
	totalCount, err := config.ArticleCollection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, false, fmt.Errorf("could not count articles: %v", err)
	}

	// Execute the query with the filter
	cursor, err := config.ArticleCollection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, false, fmt.Errorf("could not retrieve articles: %v", err)
	}
	defer cursor.Close(context.Background()) // Ensure cursor is closed

	// Extract the results into the articles slice
	err = cursor.All(context.Background(), &articles)
	if err != nil {
		return nil, false, fmt.Errorf("could not decode articles: %v", err)
	}

	// Determine if there is a next page
	hasNext := (skip + limit) < totalCount

	return articles, hasNext, nil
}
