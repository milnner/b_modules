package auth

import (
	"net/http"

	"github.com/milnner/b_modules/models"
	authSvc "github.com/milnner/b_modules/services/auth"
	"github.com/milnner/b_modules/tokens"
)

type UserAuthorizationMiddleware struct {
	tkz tokens.IJWTokenizator
}

func NewUserAuthorizationMiddleware(tkz tokens.IJWTokenizator) *UserAuthorizationMiddleware {
	return &UserAuthorizationMiddleware{tkz: tkz}
}
func (u *UserAuthorizationMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		if err := authSvc.
			NewAuthorizarionSvc(&user,
				r,
				u.tkz).Run(); err != nil ||
			user.Professor == 1 {

			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
