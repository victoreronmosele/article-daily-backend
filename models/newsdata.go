package models

type Result struct {
	Title       string   `json:"title"`
	Link        string   `json:"link"`
	Description string   `json:"description"`
	PubDate     string   `json:"pubDate"`
	Image       string   `json:"image"`
	Creators    []string `json:"creator"`
}

type NewsData struct {
	Status       string   `json:"status"`
	TotalResults int      `json:"totalResults"`
	Results      []Result `json:"results"`
	NextPage     int      `json:"nextPage"`
}