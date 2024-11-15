package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// UserServiceInterface defines the contract for user service operations
type UserServiceInterface interface {
	// UpdateUserArticles updates a user's articles in the user microservice
	// Parameters:
	// - clerkUserId: the ID of the user in Clerk
	// - articleId: the ID of the article being added
	// - price: the price of the article
	// - token: JWT token for authentication
	UpdateUserArticles(clerkUserId string, articleId string, price float64, token string) error
}

// userService implements UserServiceInterface
type userService struct {
	baseURL    string       // Base URL of the user microservice
	httpClient *http.Client // HTTP client for making requests
}

// NewUserService creates and returns a new instance of UserService
func NewUserService() UserServiceInterface {
	// Get the user service URL from environment variables
	baseURL := os.Getenv("USER_SERVICE_URL")
	
	// Return a new userService instance with configured HTTP client
	return &userService{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second, // Set timeout to prevent hanging requests
		},
	}
}

// UpdateUserArticles sends a PATCH request to update user's articles in the user service
func (s *userService) UpdateUserArticles(clerkUserId string, articleId string, price float64, token string) error {

	// Request body with article information
	requestBody := map[string]interface{}{
		"articleId": articleId,
		"price":    price,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	// The URL should include the user ID as it matches the route in user service
	url := fmt.Sprintf("%sapi/protected/users/%s", s.baseURL, clerkUserId)
	
	// Add debug logging for URL and body
	log.Printf("Sending PATCH request to: %s", url)
	log.Printf("Request body: %s", string(jsonBody))

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	req.Header.Set("X-Update-Type", "article_creation")

	

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Read response body for more detailed error information
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("user service returned non-OK status: %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}

// Singleton instance of the user service
var userServiceInstance UserServiceInterface

// GetUserService returns the singleton instance of UserService
// Creates a new instance if one doesn't exist
func GetUserService() UserServiceInterface {
	if userServiceInstance == nil {
		userServiceInstance = NewUserService()
	}
	return userServiceInstance
}

// SetUserService sets the singleton instance (used for testing)
func SetUserService(service UserServiceInterface) {
	userServiceInstance = service
}
