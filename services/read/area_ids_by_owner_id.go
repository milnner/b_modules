package read

import (
	"log"

	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type ReadAreaIdsByOwnerIdSvc struct {
	area    *models.Area
	areaIds *[]int
	repo    iRepository.IAreaRepository
	logger  *log.Logger
}

func NewReadAreaIdsByOwnerIdSvc(area *models.Area, areaIds *[]int, repo iRepository.IAreaRepository, logger *log.Logger) *ReadAreaIdsByOwnerIdSvc {
	return &ReadAreaIdsByOwnerIdSvc{area: area, areaIds: areaIds, repo: repo, logger: logger}
}

func (u *ReadAreaIdsByOwnerIdSvc) Run() error {
	var err error
	*u.areaIds, err = u.repo.GetAreaIdsByOwnerId(u.area)
	return err
}
