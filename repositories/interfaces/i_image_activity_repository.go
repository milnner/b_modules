package interfaces

import "github.com/milnner/b_modules/models"

type IImageActivityRepository interface {
	Insert(*models.ImageActivity) error
	Update(*models.ImageActivity) error
	Delete(*models.ImageActivity) error
	GetImageActivityById(*models.ImageActivity) error
	GetImageActivitiesByIds([]models.ImageActivity) error
	GetImageActivitiesByAreaId(*models.Area) ([]models.ImageActivity, error)
	GetImageActivityIdsByAreaId(*models.Area) ([]int, error)
}
