package create

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	apptypes "github.com/milnner/b_modules/apptypes"
	errapp "github.com/milnner/b_modules/errors"
	"github.com/milnner/b_modules/hasher"
	iRepositories "github.com/milnner/b_modules/repositories/interfaces"
	createService "github.com/milnner/b_modules/services/create"
)

type CreateUserController struct {
	*log.Logger
	userRepo iRepositories.IUserRepository
}

func NewCreateUserController(userRepo iRepositories.IUserRepository, logger *log.Logger) *CreateUserController {
	return &CreateUserController{userRepo: userRepo, Logger: logger}
}

func (u *CreateUserController) Handler(w http.ResponseWriter, r *http.Request) {
	var err error
	var signUpUser apptypes.SignUpUser

	if err = json.NewDecoder(r.Body).Decode(&signUpUser); err != nil {
		http.Error(w, errapp.NewJSONFormatError().Error(), http.StatusBadRequest)
		return
	}

	signUpUser.EntryDate = time.Now().String()[:19]

	createUserSvc := createService.NewCreateUserSvc(signUpUser, u.userRepo, hasher.NewBcryptHasher(), u.Logger)

	err = createUserSvc.Run()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
