package repository

import (
	"context"
	"trocup-article/config"
	"trocup-article/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetArticleByID(id string) (*models.Article, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var article models.Article
	err = config.ArticleCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&article)
	if err != nil {
		return nil, err
	}
	return &article, nil
}
