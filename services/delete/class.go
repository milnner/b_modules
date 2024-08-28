package delete

import (
	"log"

	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type DeleteClassSvc struct {
	class  *models.Class
	repo   iRepository.IClassRepository
	logger *log.Logger
}

func NewDeleteClassSvc(class *models.Class, repo iRepository.IClassRepository, logger *log.Logger) *DeleteClassSvc {
	return &DeleteClassSvc{class: class, repo: repo, logger: logger}
}

func (u *DeleteClassSvc) Run() error {
	return u.repo.Delete(u.class)
}
