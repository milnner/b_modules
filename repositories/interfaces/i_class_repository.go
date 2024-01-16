package interfaces

import "github.com/milnner/b_modules/models"

type IClassRepository interface {
	GetClassById(*models.Class) error
	Insert(*models.Class) error
	AddStudentUser(*models.Class, *models.User) error
	RemoveStudentUser(*models.Class, *models.User) error
	AddContent(class *models.Class, content *models.Content) error
	RemoveContent(class *models.Class, content *models.Content) error
	UpdateContentPosition(class *models.Class, content *models.Content) error
	Update(class *models.Class) error
	GetContentIdsById(*models.Class) ([]int, error)
	GetStudentIdsByClassId(*models.Class) ([]int, error)
	Delete(*models.Class) error
}
