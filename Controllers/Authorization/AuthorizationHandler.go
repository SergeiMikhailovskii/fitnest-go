package Authorization

import (
	"TestProject/Config"
	"TestProject/Models"
	"TestProject/Models/Authorization"
	"TestProject/Models/Base"
	RegistrationModel "TestProject/Models/Registration"
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"html/template"
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

func getPassword(login string) (error, string) {
	var record RegistrationModel.PrimaryInfo
	err := Config.DB.Model(&RegistrationModel.PrimaryInfo{}).
		Where("email", login).
		First(&record).
		Error

	return err, record.Password
}

func sendForgetPasswordEmail(to string, password string) {
	config := readEmailConfig()

	from := config.Email
	senderPassword := config.Password

	smtpHost := config.SmtpHost
	smtpPort := config.SmtpPort

	auth := smtp.PlainAuth("", from, senderPassword, smtpHost)

	t, _ := template.ParseFiles("Templates/ForgetPassword.html")
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Forget Password \n%s\n\n", mimeHeaders)))

	err := t.Execute(&body, Authorization.ForgetPasswordTemplateFields{Password: password})
	if err != nil {
		panic(err)
	}

	go func() {
		errSendMail := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, body.Bytes())
		if errSendMail != nil {
			panic(errSendMail)
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
