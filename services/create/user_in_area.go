package create

import (
	"log"

	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type CreateUserInAreaSvc struct {
	userHasAreaAccess *models.UserHasAreaAccess
	repo              iRepository.IAreaRepository
	logger            *log.Logger
}

func NewCreateUserInAreaSvc(userHasAreaAccess *models.UserHasAreaAccess, repo iRepository.IAreaRepository, logger *log.Logger) *CreateUserInAreaSvc {
	return &CreateUserInAreaSvc{userHasAreaAccess: userHasAreaAccess, repo: repo, logger: logger}
}

func (u *CreateUserInAreaSvc) Run() error {
	return u.repo.InsertUser(&u.userHasAreaAccess.Area, &u.userHasAreaAccess.User)
}
