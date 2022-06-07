package sendnotification

import (
	"article-daily-backend/server/models"
)

type SendNotification interface {
	Send(models.Notification)
}
