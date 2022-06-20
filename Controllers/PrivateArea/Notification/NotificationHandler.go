package Notification

import (
	"TestProject/Config"
	"TestProject/Controllers/Registration"
	"TestProject/Models/PrivateArea"
	"TestProject/Models/PrivateArea/DB"
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

func getNotificationsWidget(userId int) Widgets.NotificationsWidget {
	rows, err := Config.DB.Model(&DB.Notification{}).
		Where("is_active = ? AND user_id = ?", true, userId).Rows()

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
