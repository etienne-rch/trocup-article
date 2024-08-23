package services

import (
	"trocup-article/models"
	"trocup-article/repository"
)

func UpdateArticle(id string, article *models.Article) (*models.Article, error) {
	return repository.UpdateArticle(id, article)
}
