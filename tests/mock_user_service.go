package tests

import "trocup-article/services"

type MockUserService struct{}

// Update the method signature to match the interface
func (m *MockUserService) UpdateUserArticles(clerkUserId, articleId string, price float64, token string) ([]services.TransactionData, error) {
	// Return a mock response
	return []services.TransactionData{
		{
			ArticleID: articleId,
			Price:     price,
		},
	}, nil
}
