package services

import (
	"trocup-article/models"
	"trocup-article/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateArticle(article *models.Article) error {
    return repository.CreateArticle(article)
}

func GetArticles() ([]models.Article, error) {
    return repository.GetArticles()
}

func GetArticleByID(id primitive.ObjectID) (*models.Article, error) {
    return repository.GetArticleByID(id)
}