package usecases

import (
	"errors"
	"net/mail"

	"github.com/abe16s/Online-Marketplace-API/models"
	"github.com/abe16s/Online-Marketplace-API/usecases/interfaces"
	"github.com/nbutton23/zxcvbn-go"
)

type AuthUseCase struct {
    UserRepository interfaces.IUserRepo
	PwdService   interfaces.IPasswordService
	JwtService	interfaces.IJwtService
	EmailService interfaces.IEmailService
}

func (u *AuthUseCase) Register(user *models.User) (*models.User, error) {
	result := zxcvbn.PasswordStrength(user.Password, nil)
	if result.Score < 3 {
		return nil, errors.New("password is too weak")
	}
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return nil, errors.New("invalid email address")
	}
	hashedPassword, err := u.PwdService.HashPassword(user.Password)
    if err != nil {
        return nil, err
    }
    user.Password = hashedPassword

	usr, err := u.UserRepository.CreateUser(user)

	if err != nil {
		return nil, err
	}

	u.SendActivationEmail(user.Email, user.IsSeller)
	return usr, nil
}

func (u *AuthUseCase) Login(user *models.User) error {
    existingUser, err := u.UserRepository.FindByEmail(user.Email)
    if err != nil {
        return err
    }

	match := u.PwdService.ComparePassword(existingUser.Password, user.Password)
	if !match {
		return errors.New("incorrect password")
	}
	u.SendActivationEmail(user.Email, user.IsSeller)
	return nil
}

func (u *AuthUseCase) RefreshToken(email string, isSeller bool) (string, string, error) {
	newAccessToken, newRefreshToken, err := u.JwtService.GenerateToken(email, isSeller)
	if err != nil {
		return "", "", errors.New("internal server error")
	}

    return newAccessToken, newRefreshToken, nil
}


func (u *AuthUseCase) SendActivationEmail(email string, isSeller bool) error {
	token, err := u.JwtService.GenerateActivationToken(email, isSeller)
	if err != nil {
		return errors.New("internal server error")
	}

	err = u.EmailService.SendActivationEmail(email, token)
	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}


func (u *AuthUseCase) ActivateAccount(token string) (string, string, error) {
	email, isSeller, err := u.JwtService.ValidateToken(token, false)
	if err != nil {
		return "", "", errors.New("invalid token")
	}

	accessToken, refreshToken, err := u.JwtService.GenerateToken(email, isSeller)
	if err != nil {
		return "", "", errors.New("internal server error")
	}

	return accessToken, refreshToken, nil
}