package read

import (
	"log"

	"github.com/milnner/b_modules/apptypes"
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
	area := models.Area{Id: u.userArea.Area.Id, OwnerId: u.userArea.User.Id}
	if err := u.repo.GetAreaById(&area); err != nil {
		return err
	}

	if area.Activated != 0 &&
		area.OwnerId == u.userArea.User.Id {
		u.userArea.User.Permision = apptypes.Permission(apptypes.UserAreaPermissions.Write())
		return nil
	}

	return u.repo.GetPermission(&u.userArea.Area, &u.userArea.User)
}
