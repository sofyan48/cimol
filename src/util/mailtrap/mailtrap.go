package mailtrap

import (
	"os"

	"gopkg.in/gomail.v2"
)

// Mailtrap ...
type Mailtrap struct {
}

// MailtrapHandler ...
func MailtrapHandler() *Mailtrap {
	return &Mailtrap{}
}

// MailtrapInterface ...
type MailtrapInterface interface {
	SendMail(to, subject, data string) error
}

// SendMail ...
func (trap *Mailtrap) SendMail(to, subject, data string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("MAILTRAP_IDENTITY"))
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", data)
	dialer := gomail.NewDialer(
		os.Getenv("MAILTRAP_HOST"),
		587,
		os.Getenv("MAILTRAP_USERNAME"),
		os.Getenv("MAILTRAP_password"),
	)
	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}
	return nil
}
