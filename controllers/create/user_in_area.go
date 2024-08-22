package create

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/milnner/b_modules/models"
	iRepositories "github.com/milnner/b_modules/repositories/interfaces"
	authSvc "github.com/milnner/b_modules/services/auth"
	createSvc "github.com/milnner/b_modules/services/create"
	"github.com/milnner/b_modules/tokens"
)

type CreateUserInAreaController struct {
	*log.Logger
	tkz      tokens.IJWTokenizator
	areaRepo iRepositories.IAreaRepository
}

func NewCreateUserInAreaController(
	logger *log.Logger,
	tkz tokens.IJWTokenizator,
	areaRepo iRepositories.IAreaRepository) *CreateUserInAreaController {
	return &CreateUserInAreaController{Logger: logger,
		tkz:      tkz,
		areaRepo: areaRepo}
}

func (u *CreateUserInAreaController) Handler(w http.ResponseWriter, r *http.Request) {
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

	if err := createSvc.
		NewCreateUserInAreaSvc(&userHasAreaAccess,
			u.areaRepo,
			u.Logger).Run(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{}"))
}
