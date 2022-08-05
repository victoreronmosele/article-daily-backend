package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"article-daily-backend/server/config"
	"article-daily-backend/server/handlers"
	"article-daily-backend/server/services/getarticles/newsdata"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file %s", err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	newsDataKey, exists := os.LookupEnv("NEWS_DATA_KEY")

	if !exists {
		log.Fatal("NEWS_DATA_KEY not set")
	}

	config := config.Config{NewsDataAPIKey: newsDataKey}

	router := gin.Default()

	getArticlesService := newsdata.NewsData{Config: config}

	getArticleHandler := handlers.GetArticle{Service: getArticlesService}

	router.GET("/article", getArticleHandler.Run)
	router.Run(":" + port)

}
