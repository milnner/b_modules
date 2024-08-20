package delete

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/milnner/b_modules/models"
	iRepositories "github.com/milnner/b_modules/repositories/interfaces"

	deleteService "github.com/milnner/b_modules/services/delete"
	"github.com/milnner/b_modules/tests/config"
)

type DeleteUserController struct {
	userRepo iRepositories.IUserRepository
	*log.Logger
}

func NewDeleteUserController(userRepo iRepositories.IUserRepository, logger *log.Logger) *DeleteUserController {
	return &DeleteUserController{Logger: logger, userRepo: userRepo}
}

func (u *DeleteUserController) Handler(w http.ResponseWriter, r *http.Request) {
	var err error
	var user models.User

	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deleteSvc := deleteService.NewDeleteUserSvc(&user, u.userRepo, config.Logger)

	if err = deleteSvc.Run(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
