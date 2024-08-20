package read

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/milnner/b_modules/database"
	errapp "github.com/milnner/b_modules/errors"
	"github.com/milnner/b_modules/models"
	"github.com/milnner/b_modules/repositories"

	readService "github.com/milnner/b_modules/services/read"
)

type ReadUserController struct {
	*database.DatabaseConn
	*log.Logger
	jwtSecretKey string
	dbDriver     string
}

func NewReadUserController(dbConn *database.DatabaseConn, logger *log.Logger, jwtSecretKey string, dbDriver string) *ReadUserController {
	return &ReadUserController{Logger: logger, jwtSecretKey: jwtSecretKey, DatabaseConn: dbConn, dbDriver: dbDriver}
}

func (u *ReadUserController) Handler(w http.ResponseWriter, r *http.Request) {
	var err error
	var user models.User

	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, errapp.NewJSONFormatError().Error(), http.StatusBadRequest)
		return
	}

	var db *sql.DB
	var repo *repositories.UserMySQLRepository
	fmt.Println(u.dbDriver)
	if err = database.InitDatabaseConn(&db, u.DatabaseConn.User.GetSelect(), u.dbDriver); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo, err = repositories.NewUserMySQLRepository(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	readSvc := readService.NewReadUserSvc(&user, repo)

	if err = readSvc.Run(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var jsonUser []byte
	if jsonUser, err = json.Marshal(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUser)
}
