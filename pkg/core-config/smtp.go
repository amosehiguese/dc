package coreconfig

import coreutils "github.com/amosehiguese/dc/pkg/core-utils"

type smtpConfig struct {
	Username string
	Password string
	Port     string
	Host     string
	From     string
}

func setSmtpConfig() *smtpConfig {
	var sm smtpConfig
	coreutils.MustMapEnv(&sm.Username, "SMTP_USERNAME")
	coreutils.MustMapEnv(&sm.Password, "SMTP_PASSWORD")
	coreutils.MustMapEnv(&sm.Port, "SMTP_PORT")
	coreutils.MustMapEnv(&sm.Host, "SMTP_HOST")
	coreutils.MustMapEnv(&sm.From, "SMTP_FROM")

	return &sm
}

// smtp.PlainAuth("", username, password, smtpost)
// from := "amsoea@gmaog"
// to := []string{"emrekw@"}
// message := []byte("to: skfjdskfjkdsfjk")
// smtpUrl := smtpHost + ":587"

// err := smtp.SendMail(smtpUrl, auth, from, to, message)
