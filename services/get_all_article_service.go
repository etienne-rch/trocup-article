package services

import (
	"trocup-article/models"
	"trocup-article/repository"
)

func GetAllArticles(skip, limit int64) ([]models.Article, error) {
	return repository.GetAllArticles(skip, limit)
}
