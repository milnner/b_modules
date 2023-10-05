package interfaces

import "github.com/milnner/b_modules/models"

type ClassRepository interface {
	GetClassById(int) (*models.Class, error)
	GetClassByCreatorUserId(int) (*models.Class, error)
	Insert(*models.Class) error
	AddStudentUser(*models.Class, *models.User) (*models.Class, error)
	AddEditorUser(*models.Class, *models.User) (*models.Class, error)
}
