package create

import (
	"log"

	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type CreateClassSvc struct {
	class  *models.Class
	repo   iRepository.IClassRepository
	logger *log.Logger
}

func NewCreateClassSvc(class *models.Class, repo iRepository.IClassRepository, logger *log.Logger) *CreateClassSvc {
	return &CreateClassSvc{class: class, repo: repo, logger: logger}
}

func (u *CreateClassSvc) Run() error {
	return u.repo.Insert(u.class)
}
