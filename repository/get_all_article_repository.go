package repository

import (
	"context"
	"trocup-article/config"
	"trocup-article/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllArticles() ([]models.Article, error) {
	var articles []models.Article
	cursor, err := config.ArticleCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &articles)
	return articles, err
}
