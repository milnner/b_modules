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

type ReadAreasController struct {
	*database.DatabaseConn
	*log.Logger
	tkz      tokens.IJWTokenizator
	areaRepo iRepositories.IAreaRepository
}

func NewReadAreasController(areaRepo iRepositories.IAreaRepository,
	logger *log.Logger,
	tkz tokens.IJWTokenizator) *ReadAreasController {
	return &ReadAreasController{
		areaRepo: areaRepo,
		Logger:   logger,
		tkz:      tkz}
}

func (u *ReadAreasController) Handler(w http.ResponseWriter, r *http.Request) {
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
	areaIds := []int{}

	if err = json.
		NewDecoder(r.Body).
		Decode(&areaIds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	areas := []models.Area{}
	for _, i := range areaIds {
		areas = append(areas, models.Area{Id: i, OwnerId: user.Id})
	}

	if err = readSvc.
		NewReadAreasSvc(areas,
			u.areaRepo,
			u.Logger).
		Run(); err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}

	if err = json.NewEncoder(w).Encode(areas); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
