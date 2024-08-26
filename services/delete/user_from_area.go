package delete

import (
	"log"

	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type DeleteUserFromAreaSvc struct {
	area   *models.Area
	user   *models.User
	repo   iRepository.IAreaRepository
	logger *log.Logger
}

func NewDeleteUserFromAreaSvc(userHasAreaAccess *models.UserHasAreaAccess, repo iRepository.IAreaRepository, logger *log.Logger) *DeleteUserFromAreaSvc {
	return &DeleteUserFromAreaSvc{area: &userHasAreaAccess.Area, user: &userHasAreaAccess.User, repo: repo, logger: logger}
}

func (u *DeleteUserFromAreaSvc) Run() error {
	return u.repo.RemoveUser(u.area, u.user)
}
