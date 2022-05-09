package Dashboard

import (
	"TestProject/Config"
	"TestProject/Models/PrivateArea/DB"
	"fmt"
	"time"
)

func GenerateNotificationsStub(userId int) {
	for i := 0; i < 10; i++ {
		notification := DB.Notification{
			UserId: userId,
			Text:   fmt.Sprintf("Test Notification %d", i),
			Date:   time.Now(),
		}
		Config.DB.Create(&notification)
	}
}
