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
			var notification Widgets.Notification
			err = Config.DB.ScanRows(rows, &notification)

			if err != nil {
				return Widgets.NotificationsWidget{Notifications: nil}
			}

			notifications = append(notifications, notification)
		}
		return Widgets.NotificationsWidget{Notifications: notifications}
	}
}
