package create

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/models"
	iRepositories "github.com/milnner/b_modules/repositories/interfaces"

	authSvc "github.com/milnner/b_modules/services/auth"
	createSvc "github.com/milnner/b_modules/services/create"
	"github.com/milnner/b_modules/tokens"
)

type CreateAreaController struct {
	*database.DatabaseConn
	*log.Logger
	dbDriver string
	tkz      tokens.IJWTokenizator
	areaRepo iRepositories.IAreaRepository
}

func NewCreateAreaController(areaRepo iRepositories.IAreaRepository, logger *log.Logger, tkz tokens.IJWTokenizator, dbDriver string) *CreateAreaController {
	return &CreateAreaController{dbDriver: dbDriver, areaRepo: areaRepo, Logger: logger, tkz: tkz}
}

func (u *CreateAreaController) Handler(w http.ResponseWriter, r *http.Request) {
	token := tokens.ExtractTokenFromRequest(r)
	if token == "" {
		http.Error(w, "Missing Token", http.StatusUnauthorized)
		return
	}

	user := models.User{}
	authorizationSvc := authSvc.NewAuthorizarionSvc(&user, token, u.tkz)

	if err := authorizationSvc.Run(); err != nil || user.Professor == 1 {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var err error
	area := models.Area{}

	if err = json.NewDecoder(r.Body).Decode(&area); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	area.OwnerId = user.Id

	createClassSvc := createSvc.NewCreateAreaSvc(&area, u.areaRepo, u.Logger)

	if err = createClassSvc.Run(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var jsonArea []byte
	if jsonArea, err = json.Marshal(area); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonArea)
}
