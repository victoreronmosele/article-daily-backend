package models

type Article struct {
	Title       string   `json:"title"`
	Link        string   `json:"link"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Creators    []string `json:"creator"`
}