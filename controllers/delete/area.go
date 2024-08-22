package delete

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/models"
	iRepositories "github.com/milnner/b_modules/repositories/interfaces"
	authSvc "github.com/milnner/b_modules/services/auth"
	deleteSvc "github.com/milnner/b_modules/services/delete"
	"github.com/milnner/b_modules/tokens"
)

type DeleteAreaController struct {
	*database.DatabaseConn
	*log.Logger
	dbDriver string
	tkz      tokens.IJWTokenizator
	areaRepo iRepositories.IAreaRepository
}

func NewDeleteAreaController(areaRepo iRepositories.IAreaRepository,
	logger *log.Logger,
	tkz tokens.IJWTokenizator,
	dbDriver string) *DeleteAreaController {
	return &DeleteAreaController{dbDriver: dbDriver,
		areaRepo: areaRepo,
		Logger:   logger,
		tkz:      tkz}
}

func (u *DeleteAreaController) Handler(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	if err := authSvc.
		NewAuthorizarionSvc(&user,
			r,
			u.tkz).Run(); err != nil ||
		user.Professor == 1 {

		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	area := models.Area{}

	if err := json.
		NewDecoder(r.Body).
		Decode(&area); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	area.OwnerId = user.Id

	if err := deleteSvc.
		NewDeleteAreaSvc(&area,
			u.areaRepo,
			u.Logger).
		Run(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
