package repository

import (
	"context"
	"trocup-article/config"
	"trocup-article/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateArticle(article *models.Article) (*models.Article, error) {
	result, err := config.ArticleCollection.InsertOne(context.Background(), article)
	if err != nil {
		return nil, err
	}
	article.ID = result.InsertedID.(primitive.ObjectID)
	return article, nil
}
