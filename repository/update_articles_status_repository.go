package repository

import (
	"context"
	"errors"
	"time"
	"trocup-article/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleUpdateResponse struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
}

func UpdateArticlesStatus(articleIDs []primitive.ObjectID, status string) ([]ArticleUpdateResponse, error) {
	// First, get the articles with their prices
	filter := bson.M{"_id": bson.M{"$in": articleIDs}}
	cursor, err := config.ArticleCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var articles []struct {
		ID    primitive.ObjectID `bson:"_id"`
		Price float64           `bson:"price"`
	}
	if err = cursor.All(context.Background(), &articles); err != nil {
		return nil, err
	}

	// Perform the status update
	update := bson.M{"$set": bson.M{
		"status":       status,
		"lastModified": primitive.NewDateTimeFromTime(time.Now().UTC()),
	}}

	result, err := config.ArticleCollection.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("no articles found")
	}

	if result.MatchedCount != result.ModifiedCount {
		return nil, errors.New("some articles could not be updated")
	}

	// Convert to response format
	response := make([]ArticleUpdateResponse, len(articles))
	for i, article := range articles {
		response[i] = ArticleUpdateResponse{
			ID:    article.ID.Hex(),
			Price: article.Price,
		}
	}

	return response, nil
} 