package services

type MockUserService struct{}

func (m *MockUserService) UpdateUserArticles(clerkUserId, articleId string, price float64, token string) error {
    return nil
} 