package services

import (
	"trocup-article/models"
	"trocup-article/repository"
)

func CreateArticle(article *models.Article) (*models.Article, error) {
	return repository.CreateArticle(article)
}
