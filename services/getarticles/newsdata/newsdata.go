package newsdata

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"article-daily-backend/server/config"
	"article-daily-backend/server/models"
)

type NewsData struct {
	Config config.Config
}

func (n NewsData) Fetch() ([]models.Article, error) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	startTime := time.Now()

	var articles []models.Article

	res, err := http.Get("https://newsdata.io/api/1/news?apikey=" + n.Config.NewsDataAPIKey + "&language=en")

	if err != nil {
		return []models.Article{}, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	log.Info().Msgf("Response body: %s", body)

	if err != nil {
		return []models.Article{}, err
	}

	var newsData models.NewsData
	if err := json.Unmarshal(body, &newsData); err != nil {
		return []models.Article{}, err
	}

	for i, result := range newsData.Results {
		article := models.Article{
			Title:       result.Title,
			Link:        result.Link,
			Description: result.Description,
			Image:       result.Image,
			Creators:    result.Creators,
		}
		articles = append(articles, article)

		// Log summary of each article
		log.Printf("Article %d: Title: %s, Link: %s", i+1, article.Title, article.Link)
	}

	duration := time.Since(startTime)
	log.Printf("Finished fetching articles. Took %v", duration)

	return articles, nil
}
