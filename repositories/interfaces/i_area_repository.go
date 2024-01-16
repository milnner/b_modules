package interfaces

import "github.com/milnner/b_modules/models"

type IAreaRepository interface {
	GetAll() ([]models.Area, error)
	GetAreaById(*models.Area) error
	GetAreasByUserId(*models.User) ([]models.Area, error)
	GetAreaClassIdsById(*models.Class) ([]int, error)
	GetAreaContentIdsById(*models.Content) ([]int, error)
	GetAreasByIds([]models.Area) error
	Insert(*models.Area) error
	Update(*models.Area) error
	Delete(*models.Area) error
}
