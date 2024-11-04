package sendnotification

import (
	"article-daily-backend/pkg/models"
)

type SendNotification interface {
	SendNotification(models.Notification)
}
