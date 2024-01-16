package interfaces

import "github.com/milnner/b_modules/models"

type IUserRepository interface {
	GetUserById(int) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	GetAll() ([]models.User, error)
	GetUsersByIds([]int) ([]models.User, error)
	Update(*models.User) error
	Insert(*models.User) error
	Delete(*models.User) error
}
