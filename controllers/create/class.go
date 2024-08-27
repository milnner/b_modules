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

type CreateClassController struct {
	logger    *log.Logger
	classRepo iRepositories.IClassRepository
	areaRepo  iRepositories.IAreaRepository
	tkz       tokens.IJWTokenizator
}

func NewCreateClassController(classRepo iRepositories.IClassRepository,
	areaRepo iRepositories.IAreaRepository,
	tkz tokens.IJWTokenizator,
	logger *log.Logger) *CreateClassController {
	return &CreateClassController{
		classRepo: classRepo,
		areaRepo:  areaRepo,
		tkz:       tkz,
		logger:    logger,
	}
}

func (u *CreateClassController) Handler(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	if err := authSvc.NewAuthorizarionSvc(&user,
		r,
		u.tkz).
		Run(); err != nil ||
		user.Professor == 1 {

		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var err error
	class := models.Class{}

	if err = json.
		NewDecoder(r.Body).
		Decode(&class); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userHasAreaAccess := models.
		UserHasAreaAccess{User: user,
		Area: models.Area{Id: class.AreaId}}

	if err = readSvc.
		NewReadAreaPermissionSvc(&userHasAreaAccess,
			u.areaRepo,
			u.logger).
		Run(); err != nil ||
		userHasAreaAccess.User.Permision !=
			apptypes.Permission(apptypes.UserAreaPermissions.Write()) {

		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	class.UserCreatorId = user.Id
	if err = createSvc.
		NewCreateClassSvc(&class,
			u.classRepo,
			u.logger).
		Run(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{}"))
}
