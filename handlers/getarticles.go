package handlers

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/victoreronmosele/article-daily-backend/services/getarticles"
)

type GetArticle struct {
	Service getarticles.GetArticles
}

func (g GetArticle) Run(context *gin.Context) {

	articles, err := g.Service.Fetch()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error retrieving articles. Please try again.",
		})

		return
	}

	articlesNumber := len(articles)

	randomArticle := articles[rand.Intn(articlesNumber-1)]

	context.JSON(http.StatusOK, gin.H{
		"data": randomArticle,
	})
}
