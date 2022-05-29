package services

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"gopkg.in/h2non/gock.v1"

	"article-daily-backend/server/models"
)

func TestGetArticles(t *testing.T) {
	//Remove this after extracting the client from the package
	newsDataKey := os.Getenv("NEWS_DATA_KEY")

	defer gock.Off()

	testJsonResponse := map[string]interface{}{
		"status":       "success",
		"totalResults": 71057,
		"results": []map[string]interface{}{
			{
				"title": "Martijn Fischer denkt dat dit zijn laatste jaar als Hazes wordt",
				"link":  "https://www.nu.nl/cultuur-overig/6182028/martijn-fischer-denkt-dat-dit-zijn-laatste-jaar-als-hazes-wordt.html",
				"keywords": []string{
					"Boek & Cultuur",
					"Media en Cultuur",
				},
				"creator": []string{
					"NU.nl",
				},
				"video_url":   "null",
				"description": "Aloy kalandjának folytatásában úgy tűnik nem sok alkalmunk lesz tétlenül, malmozva ücsörögni.",
				"content":     "null",
				"pubDate":     "2022-02-04 07:07:01",
				"image_url":   "http://www.pcguru.hu/uploads/news/mid/horizon-forbidden-west-uj-elozetesen-a-nyugati-videkek-kihivasai-hirek-4acebfe659a8326cb803-mid.jpg",
				"source_id":   "pcguru",
				"country": []string{
					"hungary",
				},
				"category": []string{
					"entertainment",
				},
				"language": "hungarian",
			},
		},
	}

	gock.New("https://newsdata.io/api/1/news?apikey=" + newsDataKey + "&language=en").
		Get("").
		Reply(200).
		JSON(testJsonResponse)

	want := []models.Article{
		{
			Title:       "Martijn Fischer denkt dat dit zijn laatste jaar als Hazes wordt",
			Link:        "https://www.nu.nl/cultuur-overig/6182028/martijn-fischer-denkt-dat-dit-zijn-laatste-jaar-als-hazes-wordt.html",
			Creators:    []string{"NU.nl"},
			Description: "Aloy kalandjának folytatásában úgy tűnik nem sok alkalmunk lesz tétlenül, malmozva ücsörögni.",
			Image:       "http://www.pcguru.hu/uploads/news/mid/horizon-forbidden-west-uj-elozetesen-a-nyugati-videkek-kihivasai-hirek-4acebfe659a8326cb803-mid.jpg",
		},
	}

	articles := GetArticles()

	if !cmp.Equal(articles, want) {
		t.Errorf("Expected %#v, got %#v", want, articles)
	}

}
