package handlers

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"

	"article-daily-backend/server/services"
)

func GetArticle(context *gin.Context) {

	articles, err := services.GetArticles()

	if (err != nil) {
		log.Fatal(err)
	}

	articlesNumber := len(articles)

	randomArticle := articles[rand.Intn(articlesNumber-1)]

	context.JSON(http.StatusOK, gin.H{
		"data": randomArticle,
	})
}
