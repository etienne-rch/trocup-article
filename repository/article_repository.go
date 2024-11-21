package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"trocup-article/config"
)

type ArticleRepository struct {
	// Add any necessary fields here
}

func (r *ArticleRepository) UpdateArticleStatus(ctx context.Context, articleID primitive.ObjectID, status string) error {
	filter := bson.M{"_id": articleID}
	update := bson.M{"$set": bson.M{"status": status}}

	result, err := config.ArticleCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("article not found")
	}

	return nil
} 