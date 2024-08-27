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

type ReadAreaIdsByOwnerIdController struct {
	*database.DatabaseConn
	*log.Logger
	tkz      tokens.IJWTokenizator
	areaRepo iRepositories.IAreaRepository
}

func NewReadAreaIdsByOwnerIdController(areaRepo iRepositories.IAreaRepository,
	logger *log.Logger,
	tkz tokens.IJWTokenizator) *ReadAreaIdsByOwnerIdController {
	return &ReadAreaIdsByOwnerIdController{
		areaRepo: areaRepo,
		Logger:   logger,
		tkz:      tkz}
}

func (u *ReadAreaIdsByOwnerIdController) Handler(w http.ResponseWriter, r *http.Request) {
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
	var areaIds []int
	area := models.Area{OwnerId: user.Id}
	if err = readSvc.
		NewReadAreaIdsByOwnerIdSvc(&area,
			&areaIds,
			u.areaRepo,
			u.Logger).
		Run(); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if len(areaIds) == 0 {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	if err = json.
		NewEncoder(w).
		Encode(areaIds); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
