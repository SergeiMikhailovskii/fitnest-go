package Authorization

import (
	"TestProject/Config"
	"TestProject/Models"
	"TestProject/Models/Authorization"
	"TestProject/Models/Base"
	RegistrationModel "TestProject/Models/Registration"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/smtp"
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
			return nil, &record.UserID
		}
	}
}

func sendEmail(to []string) {
	config := readEmailConfig()

	from := config.Email
	password := config.Password
	message := []byte("Email test")

	smtpHost := config.SmtpHost
	smtpPort := config.SmtpPort

	auth := smtp.PlainAuth("", from, password, smtpHost)

	go func() {
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
		if err != nil {
			panic(err)
		}
	}()
}

func readEmailConfig() Models.EmailConfig {
	var config Models.EmailConfig

	file, err := ioutil.ReadFile("email_config.yaml")

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	return config
}
