package repository

import (
	"context"
	"trocup-article/config"
	"trocup-article/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllArticles(skip, limit int64) ([]models.Article, error) {
	var articles []models.Article

	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)

	cursor, err := config.ArticleCollection.Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	// Extraire les résultats dans la variable articles
	err = cursor.All(context.Background(), &articles)
	if err != nil {
		return nil, err
	}

	return articles, nil
}
