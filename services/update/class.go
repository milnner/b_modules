package update

import (
	"log"

	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type UpdateClassSvc struct {
	class  *models.Class
	repo   iRepository.IClassRepository
	logger *log.Logger
}

func NewUpdateClassSvc(class *models.Class, repo iRepository.IClassRepository, logger *log.Logger) *UpdateClassSvc {
	return &UpdateClassSvc{class: class, repo: repo, logger: logger}
}

func (u *UpdateClassSvc) Run() error {
	return u.repo.Update(u.class)
}
