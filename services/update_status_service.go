package services

import (
	"fmt"
	"trocup-article/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateArticlesStatus(articleIDs []string, status string) ([]repository.ArticleUpdateResponse, error) {
	// Convert string IDs to ObjectIDs
	objectIDs := make([]primitive.ObjectID, len(articleIDs))
	
	for i, id := range articleIDs {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, fmt.Errorf("invalid article ID format: %s", id)
		}
		objectIDs[i] = objectID
	}

	// Update the articles and get the response
	response, err := repository.UpdateArticlesStatus(objectIDs, status)
	if err != nil {
		return nil, err
	}

	return response, nil
} 