package repository

import (
	"context"
	"trocup-article/config"
	"trocup-article/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ArticleCollection *mongo.Collection

func InitArticleRepository() {
    ArticleCollection = config.Client.Database("article_dev").Collection("article")
}

func CreateArticle(article *models.Article) error {
    _, err := ArticleCollection.InsertOne(context.TODO(), article)
    return err
}

func GetArticles() ([]models.Article, error) {
    var articles []models.Article
    cursor, err := ArticleCollection.Find(context.TODO(), primitive.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var article models.Article
        if err := cursor.Decode(&article); err != nil {
            return nil, err
        }
        articles = append(articles, article)
    }
    return articles, nil
}

func GetArticleByID(id primitive.ObjectID) (*models.Article, error) {
    var article models.Article
    err := ArticleCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&article)
    if err != nil {
        return nil, err
    }
    return &article, nil
}