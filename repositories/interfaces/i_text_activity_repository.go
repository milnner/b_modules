package interfaces

import (
	"github.com/milnner/b_modules/models"
)

type ITextActivityRepository interface {
	Insert(*models.TextActivity) error
	Update(*models.TextActivity) error
	Delete(*models.TextActivity) error
	GetTextActivityById(*models.TextActivity) error
	GetTextActivitiesByIds([]models.TextActivity) error
	GetTextActivitiesByAreaId(*models.Area) ([]models.TextActivity, error)
	GetTextActivityIdsByAreaId(*models.Area) ([]int, error)
}
