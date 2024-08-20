package create

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	apptypes "github.com/milnner/b_modules/apptypes"
	"github.com/milnner/b_modules/database"
	errapp "github.com/milnner/b_modules/errors"
	"github.com/milnner/b_modules/hasher"
	repositories "github.com/milnner/b_modules/repositories"
	createService "github.com/milnner/b_modules/services/create"
)

type CreateUserController struct {
	*database.DatabaseConn
	*log.Logger
	DatabaseDriver string
}

func NewCreateUserController(dbConn *database.DatabaseConn, logger *log.Logger, databaseDriver string) *CreateUserController {
	return &CreateUserController{DatabaseDriver: databaseDriver, DatabaseConn: dbConn, Logger: logger}
}

func (u *CreateUserController) Handler(w http.ResponseWriter, r *http.Request) {
	var err error
	var signUpUser apptypes.SignUpUser

	if err = json.NewDecoder(r.Body).Decode(&signUpUser); err != nil {
		http.Error(w, errapp.NewJSONFormatError().Error(), http.StatusBadRequest)
		return
	}

	signUpUser.EntryDate = time.Now().String()[:19]

	var dbConn *sql.DB

	if err = database.InitDatabaseConn(&dbConn, u.DatabaseConn.User.GetInsert(), u.DatabaseDriver); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo, err := repositories.NewUserMySQLRepository(dbConn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createUserSvc := createService.NewCreateUserSvc(signUpUser, repo, hasher.NewBcryptHasher(), u.Logger)

	err = createUserSvc.Run()

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
