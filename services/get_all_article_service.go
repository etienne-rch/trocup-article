package services

import (
	"trocup-article/models"
	"trocup-article/repository"
)

func GetAllArticles() ([]models.Article, error) {
	return repository.GetAllArticles()
}
