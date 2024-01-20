package interfaces

import "github.com/milnner/b_modules/models"

type IUserRepository interface {
	GetUserById(*models.User) error
	GetUserByEmail(*models.User) error
	GetUsersByIds([]models.User) error
	Update(*models.User) error
	Insert(*models.User) error
	Delete(*models.User) error
}
