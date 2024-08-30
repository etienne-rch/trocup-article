package repository

import (
	"context"
	"trocup-article/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteArticle(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = config.ArticleCollection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}
