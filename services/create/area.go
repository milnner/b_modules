package create

import (
	"log"

	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type CreateAreaSvc struct {
	area   *models.Area
	repo   iRepository.IAreaRepository
	logger *log.Logger
}

func NewCreateAreaSvc(area *models.Area, repo iRepository.IAreaRepository, logger *log.Logger) *CreateAreaSvc {
	return &CreateAreaSvc{area: area, repo: repo, logger: logger}
}

func (u *CreateAreaSvc) Run() error {
	return u.repo.Insert(u.area)
}
