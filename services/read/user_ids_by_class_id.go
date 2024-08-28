package read

import (
	"log"

	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type ReadUserIdsByClassIdSvc struct {
	class  *models.Class
	repo   iRepository.IClassRepository
	logger *log.Logger
}

func NewReadUserIdsByClassIdSvc(class *models.Class, repo iRepository.IClassRepository, logger *log.Logger) *ReadUserIdsByClassIdSvc {
	return &ReadUserIdsByClassIdSvc{class: class, repo: repo, logger: logger}
}

func (u *ReadUserIdsByClassIdSvc) Run() ([]int, error) {
	return u.repo.GetStudentIdsById(u.class)
}
