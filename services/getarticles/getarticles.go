package getarticles

import "github.com/victoreronmosele/article-daily-backend/models"

type GetArticles interface {
	Fetch() ([]models.Article, error)
}

