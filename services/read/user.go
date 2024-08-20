package read

import (
	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type ReadUserSvc struct {
	user           *models.User
	userRepository iRepository.IUserRepository
}

func NewReadUserSvc(user *models.User, repo iRepository.IUserRepository) *ReadUserSvc {
	return &ReadUserSvc{user: user, userRepository: repo}
}

func (u *ReadUserSvc) Run() error {
	return u.userRepository.GetUserByEmail(u.user)
}
