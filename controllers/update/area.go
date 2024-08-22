package update

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/models"
	iRepositories "github.com/milnner/b_modules/repositories/interfaces"
	authSvc "github.com/milnner/b_modules/services/auth"
	updateSvc "github.com/milnner/b_modules/services/update"
	"github.com/milnner/b_modules/tokens"
)

type UpdateAreaController struct {
	*database.DatabaseConn
	*log.Logger
	tkz      tokens.IJWTokenizator
	areaRepo iRepositories.IAreaRepository
}

func NewUpdateAreaController(areaRepo iRepositories.IAreaRepository,
	logger *log.Logger,
	tkz tokens.IJWTokenizator) *UpdateAreaController {
	return &UpdateAreaController{
		areaRepo: areaRepo,
		Logger:   logger,
		tkz:      tkz}
}

func (u *UpdateAreaController) Handler(w http.ResponseWriter, r *http.Request) {
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

	if err = updateSvc.
		NewUpdateAreaSvc(&area,
			u.areaRepo,
			u.Logger).
		Run(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{}"))
}
