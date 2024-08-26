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

type ReadAreaPermissionController struct {
	*database.DatabaseConn
	*log.Logger
	tkz      tokens.IJWTokenizator
	areaRepo iRepositories.IAreaRepository
}

func NewReadAreaPermissionController(areaRepo iRepositories.IAreaRepository,
	logger *log.Logger,
	tkz tokens.IJWTokenizator) *ReadAreaPermissionController {
	return &ReadAreaPermissionController{
		areaRepo: areaRepo,
		Logger:   logger,
		tkz:      tkz}
}

func (u *ReadAreaPermissionController) Handler(w http.ResponseWriter, r *http.Request) {
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

	userHasAreaAccess := models.UserHasAreaAccess{}

	if err = json.
		NewDecoder(r.Body).
		Decode(&userHasAreaAccess); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userHasAreaAccess.User.Permision = ""
	userHasAreaAccess.
		Area.OwnerId = userHasAreaAccess.
		User.Id

	if err = readSvc.
		NewReadAreaPermissionSvc(&userHasAreaAccess,
			u.areaRepo,
			u.Logger).
		Run(); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if userHasAreaAccess.User.Permision == "" {
		http.Error(w, "", http.StatusNotFound)
	}

	if err = json.NewEncoder(w).
		Encode(userHasAreaAccess.User); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
