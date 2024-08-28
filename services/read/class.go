package read

import (
	"log"

	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type ReadClassSvc struct {
	class  *models.Class
	repo   iRepository.IClassRepository
	logger *log.Logger
}

func NewReadClassSvc(class *models.Class, repo iRepository.IClassRepository, logger *log.Logger) *ReadClassSvc {
	return &ReadClassSvc{class: class, repo: repo, logger: logger}
}

func (u *ReadClassSvc) Run() error {
	return u.repo.GetClassById(u.class)
}
