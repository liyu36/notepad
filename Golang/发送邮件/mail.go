package mail

import (
	"crypto/tls"

	"github.com/go-gomail/gomail"
)

type Email struct {
	ServerHost string
	ServerPort int
	FromEmail  string
	FromPasswd string
	Toers      []string
	CCers      []string
}

func NewEmail(serverHost string, serverPort int, fromEmail, fromPassword string, Toers, CCers []string) *Email {
	return &Email{
		ServerHost: serverHost,
		ServerPort: serverPort,
		FromEmail:  fromEmail,
		FromPasswd: fromPassword,
		Toers:      Toers,
		CCers:      CCers,
	}
}

func (e *Email) setHeader(m *gomail.Message, subject string) {
	if len(e.Toers) == 0 {
		return
	}
	m.SetHeader("To", e.Toers...)
	if len(e.CCers) != 0 {
		m.SetHeader("Cc", e.CCers...)
	}
	m.SetAddressHeader("From", e.FromEmail, "")
	m.SetHeader("Subject", subject)
}

func (e *Email) Send(subject, body string) error {
	m := gomail.NewMessage()
	e.setHeader(m, subject)
	m.SetBody("text/html", body)
	c := gomail.NewDialer(e.ServerHost, e.ServerPort, e.FromEmail, e.FromPasswd)
	c.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return c.DialAndSend(m)
}
