package delete

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/milnner/b_modules/models"
	iRepositories "github.com/milnner/b_modules/repositories/interfaces"
	authSvc "github.com/milnner/b_modules/services/auth"
	deleteSvc "github.com/milnner/b_modules/services/delete"
	"github.com/milnner/b_modules/tokens"
)

type DeleteUserFromAreaController struct {
	*log.Logger
	tkz      tokens.IJWTokenizator
	areaRepo iRepositories.IAreaRepository
}

func NewDeleteUserFromAreaController(
	logger *log.Logger,
	tkz tokens.IJWTokenizator,
	areaRepo iRepositories.IAreaRepository) *DeleteUserFromAreaController {
	return &DeleteUserFromAreaController{Logger: logger,
		tkz:      tkz,
		areaRepo: areaRepo}
}

func (u *DeleteUserFromAreaController) Handler(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	if err := authSvc.
		NewAuthorizarionSvc(&user,
			r,
			u.tkz).Run(); err != nil ||
		user.Professor == 1 {

		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	userHasAreaAccess := models.UserHasAreaAccess{}
	json.NewDecoder(r.Body).Decode(&userHasAreaAccess)

	if err := deleteSvc.
		NewDeleteUserFromAreaSvc(&userHasAreaAccess,
			u.areaRepo,
			u.Logger).Run(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
