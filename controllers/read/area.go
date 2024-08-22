package read

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/models"
	iRepositories "github.com/milnner/b_modules/repositories/interfaces"
	authSvc "github.com/milnner/b_modules/services/auth"
	readSvc "github.com/milnner/b_modules/services/read"
	"github.com/milnner/b_modules/tokens"
)

type ReadAreaController struct {
	*database.DatabaseConn
	*log.Logger
	tkz      tokens.IJWTokenizator
	areaRepo iRepositories.IAreaRepository
}

func NewReadAreaController(areaRepo iRepositories.IAreaRepository,
	logger *log.Logger,
	tkz tokens.IJWTokenizator) *ReadAreaController {
	return &ReadAreaController{
		areaRepo: areaRepo,
		Logger:   logger,
		tkz:      tkz}
}

func (u *ReadAreaController) Handler(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	if err := authSvc.
		NewAuthorizarionSvc(&user,
			r,
			u.tkz).Run(); err != nil ||
		user.Professor == 1 {

		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var err error
	area := models.Area{}

	if err = json.
		NewDecoder(r.Body).
		Decode(&area); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	area.OwnerId = user.Id

	if err = readSvc.
		NewReadAreaSvc(&area,
			u.areaRepo,
			u.Logger).
		Run(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.NewEncoder(w).Encode(area); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(area)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}
