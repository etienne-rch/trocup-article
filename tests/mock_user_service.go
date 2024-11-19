package tests

import "your-project/services"  // adjust this import path to match your project

type MockUserService struct{}

func (m *MockUserService) UpdateUserArticles(clerkUserId, articleId string, price float64, token string) ([]services.TransactionData, error) {
    return []services.TransactionData{{
        ArticleID: articleId,
        Price:     price,
    }}, nil
} 