package fcm

type Notification struct {
	title    string
	body     string
	imageUrl string
}

func BuildNotification(title string,
	body string,
	imageUrl string) Notification {
	return Notification{
		title:    title,
		body:     body,
		imageUrl: imageUrl,
	}
}

func (n Notification) Title() string {
	return n.title
}

func (n Notification) Body() string {
	return n.body
}

func (n Notification) ImageUrl() string {
	return n.imageUrl
}
