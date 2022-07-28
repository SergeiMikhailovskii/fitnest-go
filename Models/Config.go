package Models

type EmailConfig struct {
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
	SmtpHost string `yaml:"smtpHost"`
	SmtpPort string `yaml:"smtpPort"`
}
