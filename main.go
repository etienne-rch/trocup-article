package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"trocup-article/config"
	"trocup-article/handlers"
	"trocup-article/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    app := fiber.New()

    // Initialize MongoDB
    config.InitMongo()

    // Initialize the article repository
    repository.InitArticleRepository()

    // Set up routes
    handlers.SetupRoutes(app)

    // Get port from environment variable
    port := os.Getenv("PORT")
    if port == "" {
        port = "5002" // Default port if not specified
    }

    // Handle graceful shutdown
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        fmt.Println("Gracefully shutting down...")
        if err := config.Client.Disconnect(context.TODO()); err != nil {
            log.Fatal(err)
        }
        os.Exit(0)
    }()

    log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
