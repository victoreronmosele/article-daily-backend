package newsdata

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"article-daily-backend/server/config"
	"article-daily-backend/server/models"
)

type NewsData struct {
	Config config.Config
}

func (n NewsData) Fetch() ([]models.Article, error) {
	var articles []models.Article

	res, err := http.Get("https://newsdata.io/api/1/news?apikey=" + n.Config.NewsDataAPIKey + "&language=en")

	if err != nil {
		return []models.Article{}, err
	}



	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)


	if err != nil {
		return []models.Article{}, err
	}


	var newsData models.NewsData
	if err := json.Unmarshal(body, &newsData); err != nil {
		return []models.Article{}, err
	}

	log.Println("3")


	results := newsData.Results

	for _, result := range results {

		newsDataItemTitle := result.Title
		newsDataItemLink := result.Link
		newsDataItemDescription := result.Description
		newsDataItemImage := result.Image
		newsDataItemCreators := result.Creators

		article := models.Article{
			Title:       newsDataItemTitle,
			Link:        newsDataItemLink,
			Description: newsDataItemDescription,
			Image:       newsDataItemImage,
			Creators:    newsDataItemCreators,
		}
		articles = append(articles, article)
	}

	return articles, nil
}
