package interfaces

import "github.com/milnner/b_modules/models"

type IAreaRepository interface {
	GetAreaById(*models.Area) error
	GetUserIdsByAreaId(*models.Area) ([]int, error)
	GetPermission(*models.Area, *models.User) error
	GetAreasByIds([]models.Area) error
	InsertUser(*models.Area, *models.User) error
	RemoveUser(*models.Area, *models.User) error
	Insert(*models.Area) error
	Update(*models.Area) error
	Delete(*models.Area) error
}
