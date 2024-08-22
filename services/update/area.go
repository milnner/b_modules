package update

import (
	"log"

	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type AreaUpdateSvc struct {
	area   *models.Area
	repo   iRepository.IAreaRepository
	logger *log.Logger
}

func NewUpdateAreaSvc(area *models.Area, repo iRepository.IAreaRepository, logger *log.Logger) *AreaUpdateSvc {
	return &AreaUpdateSvc{area: area, repo: repo, logger: logger}
}

func (u *AreaUpdateSvc) Run() error {
	return u.repo.Update(u.area)
}
