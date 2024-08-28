package read

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/milnner/b_modules/apptypes"
	"github.com/milnner/b_modules/models"
	iRepositories "github.com/milnner/b_modules/repositories/interfaces"
	authSvc "github.com/milnner/b_modules/services/auth"
	readSvc "github.com/milnner/b_modules/services/read"
	"github.com/milnner/b_modules/tokens"
)

type ReadClassIdsByAreaIdController struct {
	logger    *log.Logger
	tkz       tokens.IJWTokenizator
	classRepo iRepositories.IClassRepository
	areaRepo  iRepositories.IAreaRepository
}

func NewReadClassIdsByAreaIdController(classRepo iRepositories.IClassRepository,
	areaRepo iRepositories.IAreaRepository,
	logger *log.Logger,
	tkz tokens.IJWTokenizator) *ReadClassIdsByAreaIdController {
	return &ReadClassIdsByAreaIdController{
		classRepo: classRepo,
		areaRepo:  areaRepo,
		logger:    logger,
		tkz:       tkz}
}

func (u *ReadClassIdsByAreaIdController) Handler(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println("Controller")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userHasAreaAccess := models.
		UserHasAreaAccess{User: user,
		Area: models.Area{Id: area.Id}}

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
	class := models.Class{AreaId: area.Id}

	var classIds []int
	if err := readSvc.
		NewReadClassIdsByAreaIdSvc(&class,
			&classIds,
			u.classRepo,
			u.logger).
		Run(); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if len(classIds) == 0 {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	if err := json.
		NewEncoder(w).
		Encode(classIds); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
