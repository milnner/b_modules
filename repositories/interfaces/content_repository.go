package interfaces

import "github.com/milnner/b_modules/models"

type ContentRepository interface {
	GetContentById(int) (*models.Content, error)
	GetContentByCreatorUserId(id int) (*models.Content, error)
	Insert(*models.Content) error
}
