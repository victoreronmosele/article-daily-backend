package fcm

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"

	"article-daily-backend/server/models"
)

type FCM struct {
}

func (f FCM) SendNotification(notification models.Notification) {

	ctx := context.Background()
	opt := option.WithCredentialsFile("internal/firebase/key.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	topic := "music"

	message := &messaging.Message{
		Notification:  &messaging.Notification{
			Title: notification.Title,
			Body: notification.Body,
			ImageURL: notification.ImageUrl,
		},
		Topic: topic,
	}

	client, err := app.Messaging(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	// Send a message to the devices subscribed to the provided topic.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)

}
