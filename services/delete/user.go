package delete

import (
	"log"

	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type DeleteUserSvc struct {
	*models.User
	repo iRepository.IUserRepository
}

func NewDeleteUserSvc(user *models.User, repo iRepository.IUserRepository, logger *log.Logger) *DeleteUserSvc {
	return &DeleteUserSvc{User: user, repo: repo}
}

func (u *DeleteUserSvc) Run() error {

	err := u.repo.Delete(u.User)
	return err
}
