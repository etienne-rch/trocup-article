package repository

import (
	"context"
	"trocup-article/config"
	"trocup-article/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateArticle(id string, article *models.Article) (*models.Article, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	_, err = config.ArticleCollection.ReplaceOne(context.Background(), bson.M{"_id": objectID}, article)
	if err != nil {
		return nil, err
	}
	return article, nil
}
