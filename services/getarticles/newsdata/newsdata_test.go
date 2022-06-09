package newsdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"gopkg.in/h2non/gock.v1"

	"article-daily-backend/server/config"
	"article-daily-backend/server/models"
)

func TestFetch(t *testing.T) {
	defer gock.Off()

	mockNewsDataKey := "mock-news-data-key"

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

	gock.New("https://newsdata.io/api/1/news?apikey=" + mockNewsDataKey + "&language=en").
		Get("").
		Reply(200).
		JSON(testJsonResponse)

	expected := []models.Article{
		{
			Title:       "Martijn Fischer denkt dat dit zijn laatste jaar als Hazes wordt",
			Link:        "https://www.nu.nl/cultuur-overig/6182028/martijn-fischer-denkt-dat-dit-zijn-laatste-jaar-als-hazes-wordt.html",
			Creators:    []string{"NU.nl"},
			Description: "Aloy kalandjának folytatásában úgy tűnik nem sok alkalmunk lesz tétlenül, malmozva ücsörögni.",
			Image:       "http://www.pcguru.hu/uploads/news/mid/horizon-forbidden-west-uj-elozetesen-a-nyugati-videkek-kihivasai-hirek-4acebfe659a8326cb803-mid.jpg",
		},
	}

	config := config.Config{NewsDataAPIKey: mockNewsDataKey}

	newData := NewsData{Config: config}

	actual, _ := newData.Fetch()

	if !cmp.Equal(actual, expected) {
		t.Errorf("Expected %#v, got %#v", expected, actual)
	}
}

func TestFetchWithError(t *testing.T) {
	defer gock.Off()

	mockNewsDataKey := "mock-news-data-key"

	gock.New("https://newsdata.io/api/1/news?apikey=" + mockNewsDataKey + "&language=en").
		Get("").
		Reply(500)

	config := config.Config{NewsDataAPIKey: mockNewsDataKey}

	newData := NewsData{Config: config}

	_, err := newData.Fetch()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
