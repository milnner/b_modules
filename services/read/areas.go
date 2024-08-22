package read

import (
	"log"

	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type ReadAreasSvc struct {
	areas  []models.Area
	repo   iRepository.IAreaRepository
	logger *log.Logger
}

func NewReadAreasSvc(areas []models.Area, repo iRepository.IAreaRepository, logger *log.Logger) *ReadAreasSvc {
	return &ReadAreasSvc{areas: areas, repo: repo, logger: logger}
}

func (u *ReadAreasSvc) Run() error {
	return u.repo.GetAreasByIds(u.areas)
}
