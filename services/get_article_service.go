package services

import (
	"trocup-article/models"
	"trocup-article/repository"
)

func GetArticleByID(id string) (*models.Article, error) {
	return repository.GetArticleByID(id)
}
