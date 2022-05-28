package handlers

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"

	"article-daily-backend/server/services"
)

func GetArticle(context *gin.Context) {

	articles := services.GetArticles()

	articlesNumber := len(articles)

	randomArticle := articles[rand.Intn(articlesNumber-1)]

	context.JSON(http.StatusOK, gin.H{
		"data": randomArticle,
	})
}
