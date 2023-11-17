package interfaces

import "github.com/milnner/b_modules/models"

type UserRepository interface {
	GetUserById(int) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	GetAll() (*[]models.User, error)
	Insert(*models.User) error
	Delete(*models.User) error
}
