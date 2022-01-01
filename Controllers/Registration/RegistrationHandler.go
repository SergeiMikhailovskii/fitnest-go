package Registration

import (
	"TestProject/Controllers/Anthropometry"
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
	} else if !areSecondStepFieldsFilled(primaryRegistrationRecord) {
		return &Registration.Response{
			Step:             "STEP_COMPLETE_ACCOUNT",
			Fields:           Registration.CompleteStepModel{},
			ValidationSchema: Registration.CreateStepValidationSchema,
		}, nil
	} else {
		return nil, Util.RegistrationStepNotFound
	}
}

func submitRegistrationStep(c *gin.Context) error {
	primaryRegistrationRecord := getPrimaryRegistrationRecord(c)
	if !areFirstStepFieldsFilled(primaryRegistrationRecord) {
		return submitFirstRegistrationStep(c)
	}
	return Util.RegistrationStepNotFound
}

func submitFirstRegistrationStep(c *gin.Context) error {
	userId, err := getUserId(c)
	requestBody := Registration.CreateStepModel{}
	err = c.BindJSON(&requestBody)
	if err != nil {
		return err
	}

	err = Registration.SaveCreateAccountRegistrationRecordByUserId(userId, requestBody)
	return err
}

func areFirstStepFieldsFilled(primaryRegistrationRecord Registration.PrimaryInfo) bool {
	return primaryRegistrationRecord.FirstName != "" &&
		primaryRegistrationRecord.LastName != "" &&
		primaryRegistrationRecord.Email != "" &&
		primaryRegistrationRecord.Password != ""
}

func areSecondStepFieldsFilled(primaryRegistrationRecord Registration.PrimaryInfo) bool {
	hasAnthropometryRecord := Anthropometry.HasAnthropometryRecordByUserId(1)
	return primaryRegistrationRecord.Sex != "" &&
		primaryRegistrationRecord.BirthDate.IsZero() &&
		!hasAnthropometryRecord

}

func getPrimaryRegistrationRecord(c *gin.Context) Registration.PrimaryInfo {
	userId, _ := getUserId(c)
	primaryRegistrationRecord := Registration.PrimaryInfo{}
	_ = Registration.GetPrimaryRegistrationRecordByUserId(userId, &primaryRegistrationRecord)
	return primaryRegistrationRecord
}

func getUserId(c *gin.Context) (int, error) {
	cookie, err := c.Cookie(Base.AuthUserCookie.Name)
	if err != nil {
		return -1, err
	}

	user := Models.ParseJwt(cookie)
	return user.ID, nil
}
