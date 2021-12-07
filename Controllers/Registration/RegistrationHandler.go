package Registration

import (
	"TestProject/Models"
	"TestProject/Models/Base"
	"TestProject/Models/Registration"
	"TestProject/Util"
	"github.com/gin-gonic/gin"
)

func IsRegistrationFinished(c *gin.Context) bool {
	primaryRegistrationRecord := getPrimaryRegistrationRecord(c)
	return areFirstStepFieldsFilled(primaryRegistrationRecord) &&
		primaryRegistrationRecord.Sex != ""
}

func getRegistrationStep(c *gin.Context) (*Registration.Response, error) {
	primaryRegistrationRecord := getPrimaryRegistrationRecord(c)
	if !areFirstStepFieldsFilled(primaryRegistrationRecord) {
		return &Registration.Response{
			Step:             "STEP_CREATE_ACCOUNT",
			Fields:           Registration.CreateStepModel{},
			ValidationSchema: Registration.CreateStepValidationSchema,
		}, nil
	} else {
		return nil, Util.RegistrationStepNotFound
	}
}

func areFirstStepFieldsFilled(primaryRegistrationRecord Registration.PrimaryInfo) bool {
	return primaryRegistrationRecord.FirstName != "" &&
		primaryRegistrationRecord.LastName != "" &&
		primaryRegistrationRecord.Email != "" &&
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
