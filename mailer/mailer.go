package mailer

import (
	"bytes"
	"context"
	"crypto/tls"
	"html/template"

	"github.com/sonymuhamad/todo-mailer-service/config"
	"github.com/sonymuhamad/todo-mailer-service/dto"
	"gopkg.in/gomail.v2"
)

type Mailer struct {
	client *gomail.Dialer
}

func NewMailer(cfg config.EnvConfig) *Mailer {
	d := gomail.NewDialer(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUsername, cfg.SMTPPassword)
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: cfg.SMTPSkipInsecure,
	}

	return &Mailer{
		client: d,
	}
}

func (m *Mailer) TaskCreatedNotification(_ context.Context, param dto.CreateTaskNotificationParam) error {
	t, errParseTemplate := template.ParseFiles("templates/email/create-notification.gohtml")
	if errParseTemplate != nil {
		return errParseTemplate
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, param); err != nil {
		return err
	}

	result := tpl.String()
	message := gomail.NewMessage()
	message.SetHeader("From", param.From...)
	message.SetHeader("To", param.To...)
	message.SetHeader("Subject", param.Subject)
	message.SetBody("text/html", result)

	// Send the email to Bob, Cora and Dan.
	if err := m.client.DialAndSend(message); err != nil {
		return err
	}

	return nil

}
