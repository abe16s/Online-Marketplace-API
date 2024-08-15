package infrastructures

import (
    "gopkg.in/gomail.v2"
    "fmt"
)

type EmailService struct {
}

func (es *EmailService) SendActivationEmail(email, token string) error {
	m := gomail.NewMessage()
    m.SetHeader("From", "no-reply@onlinemarketplace.com")
    m.SetHeader("To", email)
    m.SetHeader("Subject", "Activate your account")
    m.SetBody("text/html", fmt.Sprintf(`
        <h1>Welcome to Our Online Marketplace!</h1>
        <p>Please click the following link to activate your account:</p>
        <a href="http://localhost:8080/activate?token=%s">Activate Account</a>
    `, token))

    d := gomail.NewDialer("sandbox.smtp.mailtrap.io", 587, "4fbe41bc81b422", "eea6e1b5b89cd9")

    return d.DialAndSend(m)
}