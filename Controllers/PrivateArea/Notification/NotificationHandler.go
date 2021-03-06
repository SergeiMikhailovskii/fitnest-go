package Notification

import (
	"TestProject/Config"
	"TestProject/Controllers/Registration"
	"TestProject/Models/PrivateArea"
	"TestProject/Models/PrivateArea/DB"
	"TestProject/Models/PrivateArea/Request"
	"TestProject/Models/PrivateArea/Widgets"
	"github.com/gin-gonic/gin"
)

func GetNotificationsPage(c *gin.Context) (*PrivateArea.Response, error) {
	userId, _ := Registration.GetUserId(c)

	widgetsMap := make(map[string]interface{})

	widgetsMap[PrivateArea.NotificationsWidget] = getNotificationsWidget(userId)

	return &PrivateArea.Response{
		Widgets: widgetsMap,
	}, nil
}

func DeactivateNotifications(c *gin.Context) error {
	var request []int

	err := c.BindJSON(&request)

	for _, item := range request {
		Config.DB.Model(&DB.Notification{}).Where("id = ?", item).Update("is_active", false)
	}

	return err
}

func PinNotification(c *gin.Context) error {
	var request Request.PinNotificationRequest

	err := c.BindJSON(&request)

	if err != nil {
		return err
	}

	err = Config.DB.
		Model(&DB.Notification{}).
		Where("id = ?", request.Id).
		Update("is_pinned", &request.Pin).
		Error

	return err
}

func DeleteNotification(c *gin.Context) error {
	var request Request.DeleteNotificationRequest

	err := c.BindJSON(&request)

	if err != nil {
		return err
	}

	err = Config.DB.Delete(&DB.Notification{}, request.Id).Error

	return err
}

func getNotificationsWidget(userId int) Widgets.NotificationsWidget {
	rows, err := Config.DB.Model(&DB.Notification{}).
		Where("user_id = ?", userId).Rows()

	if err != nil {
		return Widgets.NotificationsWidget{Notifications: nil}
	} else {
		var notifications []Widgets.Notification
		for rows.Next() {
			var notification DB.Notification
			err = Config.DB.ScanRows(rows, &notification)

			if err != nil {
				return Widgets.NotificationsWidget{Notifications: nil}
			}

			notifications = append(notifications, Widgets.Notification{
				ID:       notification.ID,
				Title:    notification.Text,
				Date:     notification.Date.Format("2006-01-02T15:04:05"),
				Type:     notification.Type,
				IsActive: notification.IsActive,
				IsPinned: notification.IsPinned,
			})
		}
		return Widgets.NotificationsWidget{Notifications: notifications}
	}
}
