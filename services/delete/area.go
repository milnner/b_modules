package delete

import (
	"log"

	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type DeleteAreaSvc struct {
	area   *models.Area
	repo   iRepository.IAreaRepository
	logger *log.Logger
}

func NewDeleteAreaSvc(area *models.Area, repo iRepository.IAreaRepository, logger *log.Logger) *DeleteAreaSvc {
	return &DeleteAreaSvc{area: area, repo: repo, logger: logger}
}

func (u *DeleteAreaSvc) Run() error {
	return u.repo.Delete(u.area)
}
