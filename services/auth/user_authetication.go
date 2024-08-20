package auth

import (
	"github.com/milnner/b_modules/models"
	"github.com/milnner/b_modules/tokens"
)

type AuthenticationSvc struct {
	user  *models.User
	token *string
	tkz   tokens.IJWTokenizator
}

func NewAuthenticationSvc(user *models.User, token *string, tkz tokens.IJWTokenizator) *AuthenticationSvc {
	return &AuthenticationSvc{user: user, token: token, tkz: tkz}
}

func (u *AuthenticationSvc) Run() error {
	var err error
	claims := make(map[string]interface{})
	claims["sub"] = u.user.Id
	claims["email"] = u.user.Email
	*u.token, err = u.tkz.GenerateToken(claims)
	return err
}
