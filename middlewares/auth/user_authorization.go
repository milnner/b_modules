package auth

import (
	"fmt"
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
		// Extrair o token da requisição
		token := tokens.ExtractTokenFromRequest(r)
		if token == "" {
			http.Error(w, "Invalid or Missing Token", http.StatusBadRequest)
			return
		}

		user := models.User{}
		authorizationSvc := authSvc.NewAuthorizarionSvc(&user, token, u.tkz)

		if err := authorizationSvc.Run(); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		fmt.Println("Passou")

		next.ServeHTTP(w, r)
	})
}
