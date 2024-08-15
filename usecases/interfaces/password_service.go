package interfaces

type IPasswordService interface {
	HashPassword(password string) (string, error)
	ComparePassword(existingPassword string, userPassword string) bool
}