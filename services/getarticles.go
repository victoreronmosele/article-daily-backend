package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"article-daily-backend/server/models"
)

func GetArticles() []models.Article {
	newsDataKey := os.Getenv("NEWS_DATA_KEY")

	var articles []models.Article

	res, err := http.Get("https://newsdata.io/api/1/news?apikey=" + newsDataKey + "&language=en")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var newsData models.NewsData
	if err := json.Unmarshal(body, &newsData); err != nil {
		log.Fatal(err)
	}

	results := newsData.Results 

	resultLength := len(results)
 
	for i := 0; i < resultLength-1; i++ {
		newsDataItem := results[i]
		newsDataItemTitle := newsDataItem.Title
		newsDataItemLink := newsDataItem.Link
		newsDataItemDescription := newsDataItem.Description
		newsDataItemImage := newsDataItem.Image
		newsDataItemCreators := newsDataItem.Creators

		article := models.Article{
			Title:       newsDataItemTitle,
			Link:        newsDataItemLink,
			Description: newsDataItemDescription,
			Image:       newsDataItemImage,
			Creators:    newsDataItemCreators,
		}
		articles = append(articles, article)
	}

	return articles
}
