package interfaces

import "github.com/abe16s/Online-Marketplace-API/models"

type IUserRepo interface {
	FindByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
}