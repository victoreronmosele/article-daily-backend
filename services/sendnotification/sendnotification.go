package sendnotification

import (
	"github.com/victoreronmosele/article-daily-backend/models"
)

type SendNotification interface {
	SendNotification(models.Notification)
}
