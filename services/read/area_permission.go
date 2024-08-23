package read

import (
	"log"

	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type ReadAreaPermissionSvc struct {
	userArea *models.UserHasAreaAccess
	repo     iRepository.IAreaRepository
	logger   *log.Logger
}

func NewReadAreaPermissionSvc(userArea *models.UserHasAreaAccess, repo iRepository.IAreaRepository, logger *log.Logger) *ReadAreaPermissionSvc {
	return &ReadAreaPermissionSvc{userArea: userArea, repo: repo, logger: logger}
}

func (u *ReadAreaPermissionSvc) Run() error {
	return u.repo.GetPermission(&u.userArea.Area, &u.userArea.User)
}
