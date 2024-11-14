package services

import (
	"errors"
	"trocup-article/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateArticlesStatus(articleIDs []string, status string) error {
	// Convert string ID to ObjectID
	objectIDs := make([]primitive.ObjectID, len(articleIDs))
	for i, id := range articleIDs {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return errors.New("invalid article ID format")
		}
		objectIDs[i] = objectID
	}

	// Call repository to update status
	err := repository.UpdateArticlesStatus(objectIDs, status)
	if err != nil {
		return err
	}

	return nil
} 