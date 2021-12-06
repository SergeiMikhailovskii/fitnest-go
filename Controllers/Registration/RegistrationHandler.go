package Registration

import (
	"TestProject/Models"
	"TestProject/Models/Base"
	"TestProject/Models/Registration"
	"github.com/gin-gonic/gin"
)

func IsRegistrationFinished(c *gin.Context) bool {
	primaryRegistrationRecord := getPrimaryRegistrationRecord(c)
	return primaryRegistrationRecord.FirstName != "" &&
		primaryRegistrationRecord.LastName != "" &&
		primaryRegistrationRecord.Email != "" &&
		primaryRegistrationRecord.Sex != "" &&
		primaryRegistrationRecord.Password != ""
}

func getPrimaryRegistrationRecord(c *gin.Context) Registration.PrimaryInfo {
	cookie, err := c.Cookie(Base.AuthUserCookie.Name)
	if err != nil {
		panic(err)
	}

	user := Models.ParseJwt(cookie)
	primaryRegistrationRecord := Registration.PrimaryInfo{}
	_ = Registration.GetPrimaryRegistrationRecordByUserId(user.ID, &primaryRegistrationRecord)
	return primaryRegistrationRecord
}
