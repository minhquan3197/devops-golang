package mails

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"html/template"
	"net"
	"net/smtp"
	"project-golang/configs"
	"strings"
)

//Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

type loginAuth struct {
	username, password string
}

func NewRequest(to []string, subject, body string) *Request {
	fmt.Println(to)
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown from server")
		}
	}
	return nil, nil
}

func (r *Request) SendEmail() (bool, error) {
	cfg := configs.Load()
	user := cfg.Mail.Sender
	pass := cfg.Mail.Password
	host := cfg.Mail.Host
	port := cfg.Mail.Port
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	newBody := strings.Replace(r.body, "\r", "<br/><br/>", -1)
	msg := []byte(subject + mime + "\n" + newBody)
	addr := host + ":" + port

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		println(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		println(err)
	}

	tlsconfig := &tls.Config{
		ServerName: host,
	}

	if err = c.StartTLS(tlsconfig); err != nil {
		println(err)
	}

	auth := LoginAuth(user, pass)
	fmt.Println("ok", auth)

	if err = c.Auth(auth); err != nil {
		println(err)
	}

	if err := smtp.SendMail(addr, auth, user, r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}
