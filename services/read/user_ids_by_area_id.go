package read

import (
	"log"

	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type ReadUserIdsByAreaIdSvc struct {
	area    *models.Area
	userIds *[]int
	repo    iRepository.IAreaRepository
	logger  *log.Logger
}

func NewReadUserIdsByAreaIdSvc(area *models.Area, userIds *[]int, repo iRepository.IAreaRepository, logger *log.Logger) *ReadUserIdsByAreaIdSvc {
	return &ReadUserIdsByAreaIdSvc{area: area, repo: repo, logger: logger, userIds: userIds}
}

func (u *ReadUserIdsByAreaIdSvc) Run() error {
	var err error
	*u.userIds, err = u.repo.GetUserIdsByAreaId(u.area)
	return err
}
