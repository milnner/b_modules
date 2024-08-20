package auth

import (
	"github.com/milnner/b_modules/apptypes"
	"github.com/milnner/b_modules/hasher"
	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
	readServices "github.com/milnner/b_modules/services/read"
)

type AuthenticateSvc struct {
	signInUser apptypes.SignInUser
	user       *models.User
	repo       iRepository.IUserRepository
	hasher     hasher.IHasher
}

func NewAuthenticateSvc(signInUser apptypes.SignInUser, user *models.User, repo iRepository.IUserRepository, hasher hasher.IHasher) *AuthenticateSvc {
	return &AuthenticateSvc{signInUser: signInUser, user: user, repo: repo, hasher: hasher}
}

func (u *AuthenticateSvc) Run() error {
	var err error

	u.user.Email = u.signInUser.Email

	readUserSvc := readServices.NewReadUserSvc(u.user, u.repo)

	if err = readUserSvc.Run(); err != nil {
		return err
	}

	err = u.hasher.Compare([]byte(u.user.Hash), []byte(u.signInUser.Password))

	return err
}
