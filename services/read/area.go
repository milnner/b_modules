package read

import (
	"log"

	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type ReadAreaSvc struct {
	area   *models.Area
	repo   iRepository.IAreaRepository
	logger *log.Logger
}

func NewReadAreaSvc(area *models.Area, repo iRepository.IAreaRepository, logger *log.Logger) *ReadAreaSvc {
	return &ReadAreaSvc{area: area, repo: repo, logger: logger}
}

func (u *ReadAreaSvc) Run() error {
	return u.repo.GetAreaById(u.area)
}
