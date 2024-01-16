package interfaces

import "github.com/milnner/b_modules/models"

type IContentRepository interface {
	GetContentById(*models.Content) error
	GetContentsByAreaId(*models.Area) ([]models.Content, error)
	GetContentsByIds([]models.Content) error
	Insert(*models.Content) error
	Update(*models.Content) error
	Delete(*models.Content) error
	AddActivity(*models.Content, interface{}) error
	RemoveActivity(*models.Content, interface{}) error
	UpdateActivityPosition(*models.Content, interface{}) error
	GetActivityIdsByContentId(*models.Content, interface{}) ([]int, error)
}
