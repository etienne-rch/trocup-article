package services

import (
	"trocup-article/repository"
)

func DeleteArticle(id string) error {
	return repository.DeleteArticle(id)
}
