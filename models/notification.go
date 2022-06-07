package models

type Notification interface {
	Title() string
	Body() string
	ImageUrl() string
}
