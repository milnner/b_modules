package delete

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/models"
	"github.com/milnner/b_modules/repositories"
	deleteService "github.com/milnner/b_modules/services/delete"
	"github.com/milnner/b_modules/tests/config"
)

type DeleteUserController struct {
	*database.DatabaseConn
	*log.Logger
	jwtSecretKey string
	dbDriver     string
}

func NewDeleteUserController(dbConn *database.DatabaseConn, logger *log.Logger, jwtSecretKey, dbDriver string) *DeleteUserController {
	return &DeleteUserController{dbDriver: dbDriver, DatabaseConn: dbConn, Logger: logger, jwtSecretKey: jwtSecretKey}
}

func (u *DeleteUserController) Handler(w http.ResponseWriter, r *http.Request) {
	var err error
	var user models.User

	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var db *sql.DB
	var repo *repositories.UserMySQLRepository
	if err = database.InitDatabaseConn(&db, u.DatabaseConn.User.GetSelect(), u.dbDriver); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo, err = repositories.NewUserMySQLRepository(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deleteSvc := deleteService.NewDeleteUserSvc(&user, repo, config.Logger)

	if err = deleteSvc.Run(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
