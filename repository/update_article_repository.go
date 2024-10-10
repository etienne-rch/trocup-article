package repository

import (
	"context"
	"trocup-article/config"
	"trocup-article/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpdateArticle met à jour un article partiellement dans la base de données
func UpdateArticle(id primitive.ObjectID, updates map[string]interface{}) (*models.Article, error) {
	// Créer l'objet de mise à jour avec les champs spécifiés dans "updates"
	update := bson.M{
		"$set": updates,
	}

	// Options pour retourner l'article mis à jour après modification
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	// Effectuer la mise à jour dans MongoDB et récupérer l'article mis à jour
	var updatedArticle models.Article
	err := config.ArticleCollection.FindOneAndUpdate(context.TODO(), bson.M{"_id": id}, update, opts).Decode(&updatedArticle)
	if err != nil {
		return nil, err
	}

	return &updatedArticle, nil
}
