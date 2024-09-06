package create

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/milnner/b_modules/apptypes"
	"github.com/milnner/b_modules/models"
	iRepositories "github.com/milnner/b_modules/repositories/interfaces"
	authSvc "github.com/milnner/b_modules/services/auth"
	createSvc "github.com/milnner/b_modules/services/create"
	readSvc "github.com/milnner/b_modules/services/read"
	"github.com/milnner/b_modules/tokens"
)

type CreateClassHasUserController struct {
	logger    *log.Logger
	tkz       tokens.IJWTokenizator
	classRepo iRepositories.IClassRepository
	areaRepo  iRepositories.IAreaRepository
}

func NewCreateClassHasUserController(
	classRepo iRepositories.IClassRepository,
	areaRepo iRepositories.IAreaRepository,
	tkz tokens.IJWTokenizator,
	logger *log.Logger,
) *CreateClassHasUserController {
	return &CreateClassHasUserController{
		classRepo: classRepo,
		areaRepo:  areaRepo,
		logger:    logger,
		tkz:       tkz,
	}
}

func (u *CreateClassHasUserController) Handler(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	if err := authSvc.
		NewAuthorizarionSvc(&user,
			r,
			u.tkz).Run(); err != nil ||
		user.Professor == 1 {

		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	userHasClassAccess := models.UserHasClassAccess{}

	json.NewDecoder(r.Body).Decode(&userHasClassAccess)

	userHasAreaAccess := models.UserHasAreaAccess{
		Area: models.
			Area{Id: userHasClassAccess.Class.AreaId},
		User: user}

	if err := readSvc.
		NewReadAreaPermissionSvc(&userHasAreaAccess,
			u.areaRepo,
			u.logger).
		Run(); err != nil ||
		!(!(userHasAreaAccess.User.Permision !=
			apptypes.Permission(apptypes.UserAreaPermissions.Write())) ||

			!(userHasAreaAccess.User.Permision !=
				apptypes.Permission(apptypes.UserAreaPermissions.Read()))) {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	if err := createSvc.
		NewCreateClassHasUserSvc(&userHasClassAccess,
			u.classRepo,
			u.logger).
		Run(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{}"))
}
