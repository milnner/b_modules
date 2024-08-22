package auth

import (
	"errors"
	"net/http"

	"github.com/milnner/b_modules/models"
	"github.com/milnner/b_modules/tokens"
)

type AuthorizarionService struct {
	user *models.User
	req  *http.Request
	tkz  tokens.IJWTokenizator
}

func NewAuthorizarionSvc(user *models.User, req *http.Request, tkz tokens.IJWTokenizator) *AuthorizarionService {
	return &AuthorizarionService{user: user, req: req, tkz: tkz}
}

func (u *AuthorizarionService) Run() error {
	token := tokens.ExtractTokenFromRequest(u.req)
	if token == "" {
		return errors.New("Missing Token")
	}

	claims, err := u.tkz.ValidateToken(token)

	if err != nil {
		return err
	}
	u.user.Id = (*claims)["sub"].(int)
	u.user.Email = (*claims)["email"].(string)
	return nil
}
