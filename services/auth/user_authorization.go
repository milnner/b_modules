package auth

import (
	"github.com/milnner/b_modules/models"
	"github.com/milnner/b_modules/tokens"
)

type AuthorizarionService struct {
	user  *models.User
	token string
	tkz   tokens.IJWTokenizator
}

func NewAuthorizarionSvc(user *models.User, token string, tkz tokens.IJWTokenizator) *AuthorizarionService {
	return &AuthorizarionService{user: user, token: token, tkz: tkz}
}

func (u *AuthorizarionService) Run() error {

	claims, err := u.tkz.ValidateToken(u.token)

	if err != nil {
		return err
	}
	u.user.Id = (*claims)["sub"].(int)
	u.user.Email = (*claims)["email"].(string)
	return nil
}
