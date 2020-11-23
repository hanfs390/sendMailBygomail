package tool

import (
	"errors"
	"github.com/outakujo/utils"
	"gopkg.in/gomail.v2"
)

type Email struct {
	From     string
	Host     string
	Port     int
	UserName string
	Password string
}

func (r *Email) Send(subject, htmlBody, attachFile, rename string, to ...string) error {
	l := len(to)
	if l == 0 {
		return errors.New("to can not be empty")
	}
	m := gomail.NewMessage()
	m.SetHeader("From", r.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlBody)
	if utils.FileIsExist(attachFile) {
		if rename != "" {
			setting := gomail.Rename(rename)
			m.Attach(attachFile, setting)
		} else {
			m.Attach(attachFile)
		}
	}
	d := gomail.NewDialer(r.Host, r.Port, r.UserName, r.Password)
	return d.DialAndSend(m)
}