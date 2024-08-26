package read

import (
	"encoding/json"
	"log"
	"net/http"

	errapp "github.com/milnner/b_modules/errors"
	"github.com/milnner/b_modules/models"
	iRepositories "github.com/milnner/b_modules/repositories/interfaces"

	readService "github.com/milnner/b_modules/services/read"
)

type ReadUserController struct {
	*log.Logger
	userRepo iRepositories.IUserRepository
}

func NewReadUserController(userRepo iRepositories.IUserRepository, logger *log.Logger) *ReadUserController {
	return &ReadUserController{Logger: logger, userRepo: userRepo}
}

func (u *ReadUserController) Handler(w http.ResponseWriter, r *http.Request) {
	var err error
	var user models.User

	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, errapp.NewJSONFormatError().Error(), http.StatusBadRequest)
		return
	}

	readSvc := readService.NewReadUserSvc(&user, u.userRepo)

	if err = readSvc.Run(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var jsonUser []byte
	if jsonUser, err = json.Marshal(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUser)
}
