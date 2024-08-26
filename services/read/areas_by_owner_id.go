package read

import (
	"log"

	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type ReadAreasByOwnerIdSvc struct {
	user   *models.User
	areas  *[]models.Area
	repo   iRepository.IAreaRepository
	logger *log.Logger
}

func NewReadAreasByOwnerIdSvc(user *models.User, areas *[]models.Area, repo iRepository.IAreaRepository, logger *log.Logger) *ReadAreasByOwnerIdSvc {
	return &ReadAreasByOwnerIdSvc{user: user, areas: areas, repo: repo, logger: logger}
}

func (u *ReadAreasByOwnerIdSvc) Run() error {
	u.repo.GetAreasByOwnerId(u.areas, u.user)
	return nil
}
