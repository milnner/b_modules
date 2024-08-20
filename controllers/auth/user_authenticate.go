package auth

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/milnner/b_modules/apptypes"
	"github.com/milnner/b_modules/database"
	errapp "github.com/milnner/b_modules/errors"
	"github.com/milnner/b_modules/hasher"
	"github.com/milnner/b_modules/models"
	"github.com/milnner/b_modules/repositories"
	"github.com/milnner/b_modules/services/auth"
	"github.com/milnner/b_modules/tokens"
)

type AuthenticateUserController struct {
	*database.DatabaseConn
	*log.Logger
	DatabaseDriver string
	hasher         hasher.IHasher
	tkz            tokens.IJWTokenizator
}

func NewAuthenticateUserController(dbConn *database.DatabaseConn, logger *log.Logger, tkz tokens.IJWTokenizator, hasher hasher.IHasher, databaseDriver string) *AuthenticateUserController {
	return &AuthenticateUserController{DatabaseDriver: databaseDriver, DatabaseConn: dbConn, Logger: logger, hasher: hasher, tkz: tkz}
}

func (u *AuthenticateUserController) Handler(w http.ResponseWriter, r *http.Request) {
	var err error
	var signInUser apptypes.SignInUser

	if err = json.NewDecoder(r.Body).Decode(&signInUser); err != nil {
		http.Error(w, errapp.NewJSONFormatError().Error(), http.StatusBadRequest)
		return
	}

	var dbConn *sql.DB

	if err = database.InitDatabaseConn(&dbConn, u.DatabaseConn.User.GetSelect(), u.DatabaseDriver); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo, err := repositories.NewUserMySQLRepository(dbConn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	readUserSvc := auth.NewAuthenticateSvc(signInUser, &user, repo, u.hasher)

	if err = readUserSvc.Run(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var token string
	authenticationSvc := auth.NewAuthenticationSvc(&user, &token, u.tkz)

	if err = authenticationSvc.Run(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Authorization", "Bearer "+token)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}
