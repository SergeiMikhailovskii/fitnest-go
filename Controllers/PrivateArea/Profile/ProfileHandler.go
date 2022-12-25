package Profile

import (
	"TestProject/Config"
	"TestProject/Controllers/Registration"
	"TestProject/Models/PrivateArea"
	"TestProject/Models/PrivateArea/Widgets"
	Registration2 "TestProject/Models/Registration"
	"github.com/gin-gonic/gin"
)

func GetProfilePage(c *gin.Context) (*PrivateArea.Response, error) {
	userId, _ := Registration.GetUserId(c)

	widgetsMap := make(map[string]interface{})

	widgetsMap[PrivateArea.ProfileInfoWidget] = getProfileInfoWidget(userId)

	return &PrivateArea.Response{
		Widgets: widgetsMap,
	}, nil
}

func getProfileInfoWidget(userId int) *Widgets.ProfileInfoWidget {
	var anthropometry Registration2.AnthropometryModel
	var goal Registration2.GoalModel
	var primaryInfo Widgets.ProfilePrimaryInfoQuery
	err := Config.DB.Where("user_id = ?", userId).Last(&anthropometry).Error
	if err != nil {
		return nil
	}
	err = Config.DB.Where("user_id = ?", userId).Last(&goal).Error
	if err != nil {
		return nil
	}
	err = Config.DB.Raw("SELECT first_name, last_name, date_part('year', AGE(now(), birth_date)) from registration_primary_info where user_id = ?", userId).Scan(&primaryInfo).Error
	if err != nil {
		return nil
	}
	return &Widgets.ProfileInfoWidget{
		FirstName: primaryInfo.FirstName,
		LastName:  primaryInfo.LastName,
		Program:   goal.Goal,
		Height:    int(anthropometry.Height),
		Weight:    int(anthropometry.Weight),
		Age:       primaryInfo.DatePart,
	}
}
