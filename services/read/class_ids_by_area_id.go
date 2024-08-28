package read

import (
	"log"

	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type ReadClassIdsByAreaIdSvc struct {
	class    *models.Class
	classIds *[]int
	repo     iRepository.IClassRepository
	logger   *log.Logger
}

func NewReadClassIdsByAreaIdSvc(class *models.Class, classIds *[]int, repo iRepository.IClassRepository, logger *log.Logger) *ReadClassIdsByAreaIdSvc {
	return &ReadClassIdsByAreaIdSvc{class: class, classIds: classIds, repo: repo, logger: logger}
}

func (u *ReadClassIdsByAreaIdSvc) Run() error {
	var err error
	*u.classIds, err = u.repo.GetClassIdsByAreaId(u.class)
	return err
}
