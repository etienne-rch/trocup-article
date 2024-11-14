package services

import (
	"trocup-article/models"
	"trocup-article/repository"
)

func GetAllArticles(skip, limit int64, category string) ([]models.Article, bool, error) {
	return repository.GetAllArticles(skip, limit, category)
}
