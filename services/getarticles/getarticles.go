package getarticles

import "article-daily-backend/server/models"

type GetArticles interface {
	Fetch() ([]models.Article, error)
}

