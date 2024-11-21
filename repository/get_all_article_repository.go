package repository

import (
	"context"
	"fmt"
	"trocup-article/config"
	"trocup-article/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllArticles(skip, limit int64, category string, status string) ([]models.Article, bool, error) {
	var articles []models.Article

	// Créer des options de recherche
	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)

	filter := bson.M{}
	if category != "" {
		filter["category"] = category
	}
	if status != "" {
		filter["status"] = status
	}

	// Compter le nombre total d'articles
	totalCount, err := config.ArticleCollection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, false, fmt.Errorf("could not count articles: %v", err)
	}

	// Exécuter la recherche
	cursor, err := config.ArticleCollection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, false, err
	}
	defer cursor.Close(context.Background()) // Assurez-vous de fermer le curseur

	// Extraire les résultats dans la variable articles
	err = cursor.All(context.Background(), &articles)
	if err != nil {
		return nil, false, err
	}

	// Vérifier s'il y a une page suivante
	hasNext := (skip + limit) < totalCount

	return articles, hasNext, nil
}
