package main

import (
	"log"
	"math/rand"
	"os"

	"github.com/joho/godotenv"

	"article-daily-backend/server/config"
	"article-daily-backend/server/models"
	"article-daily-backend/server/services/getarticles/newsdata"
	"article-daily-backend/server/services/sendnotification/fcm"
)

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	newsDataKey, exists := os.LookupEnv("NEWS_DATA_KEY")

	if !exists {
		log.Fatal("NEWS_DATA_KEY not set")
	}

	sendNotificationService := fcm.FCM{}
	config := config.Config{NewsDataAPIKey: newsDataKey}
	getArticlesService := newsdata.NewsData{Config: config}

	articles, err := getArticlesService.Fetch()

	if err != nil {
		log.Println(err)
		log.Fatal("Error fetching articles")
	}

	articlesNumber := len(articles)

	randomArticle := articles[rand.Intn(articlesNumber-1)]

	articleTitle := randomArticle.Title
	articleDescription := randomArticle.Description
	var articleImage string = randomArticle.Image
	if &randomArticle.Image == nil {
		articleImage = "https://images.pexels.com/photos/4711052/pexels-photo-4711052.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260"
	}

	notification := models.Notification{Title: articleTitle, Body: articleDescription, ImageUrl: articleImage}

	sendNotificationService.SendNotification(notification)

}
