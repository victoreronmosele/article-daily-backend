package sendnotification

import (
	"article-daily-backend/server/models"
)

type SendNotification interface {
	SendNotification(models.Notification)
}
