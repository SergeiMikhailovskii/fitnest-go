package Authorization

import (
	"TestProject/Config"
	"TestProject/Models/Authorization"
	"TestProject/Models/Base"
	RegistrationModel "TestProject/Models/Registration"
)

func getLoginFields() Authorization.GetLoginFields {
	return Authorization.GetLoginFields{}
}

func getLoginValidationSchema() Authorization.LoginValidationSchemaType {
	return Authorization.LoginValidationSchema
}

func loginUser(fields Authorization.GetLoginFields) (*Base.Error, *int) {
	err := Config.DB.Model(&RegistrationModel.PrimaryInfo{}).
		Where("email", fields.Login).
		First(&RegistrationModel.PrimaryInfo{}).
		Error

	if err != nil {
		return &Base.Error{
			Field:   "login",
			Message: "error.invalid",
		}, nil
	} else {
		var record RegistrationModel.PrimaryInfo
		err = Config.DB.Model(&RegistrationModel.PrimaryInfo{}).
			Where("email", fields.Login).
			Where("password", fields.Password).
			First(&record).
			Error

		if err != nil {
			return &Base.Error{
				Field:   "password",
				Message: "error.invalid",
			}, nil
		} else {
			return nil, &record.ID
		}
	}
}
