package repository

import (
	"context"
	"errors"
	"time"
	"trocup-article/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateArticlesStatus(articleIDs []primitive.ObjectID, status string) error {
	filter := bson.M{"_id": bson.M{"$in": articleIDs}}

	update := bson.M{"$set": bson.M{
		"status":       status,
		"lastModified": primitive.NewDateTimeFromTime(time.Now().UTC()),
	}}

	result, err := config.ArticleCollection.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("no articles found")
	}

	if result.MatchedCount != result.ModifiedCount {
		return errors.New("some articles could not be updated")
	}

	return nil
} 