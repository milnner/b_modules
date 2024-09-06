package create

import (
	"log"

	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type CreateClassHasUserSvc struct {
	userHasClassAccess *models.UserHasClassAccess
	repo               iRepository.IClassRepository
	logger             *log.Logger
}

func NewCreateClassHasUserSvc(userHasClassAccess *models.UserHasClassAccess, repo iRepository.IClassRepository, logger *log.Logger) *CreateClassHasUserSvc {
	return &CreateClassHasUserSvc{userHasClassAccess: userHasClassAccess, repo: repo, logger: logger}
}

func (u *CreateClassHasUserSvc) Run() error {
	return u.repo.AddStudentUser(&u.userHasClassAccess.Class, &u.userHasClassAccess.User)
}
