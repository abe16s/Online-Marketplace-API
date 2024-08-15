package interfaces

type IEmailService interface {
	SendActivationEmail(email, token string) error
}