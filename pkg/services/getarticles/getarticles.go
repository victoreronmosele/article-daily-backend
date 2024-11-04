package getarticles

import "article-daily-backend/pkg/models"

type GetArticles interface {
	Fetch() ([]models.Article, error)
}

