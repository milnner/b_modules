package create

import (
	"log"
	"time"

	"github.com/milnner/b_modules/apptypes"
	"github.com/milnner/b_modules/datacheck"
	"github.com/milnner/b_modules/hasher"
	"github.com/milnner/b_modules/models"
	iRepository "github.com/milnner/b_modules/repositories/interfaces"
)

type CreateUserSvc struct {
	user   apptypes.SignUpUser
	repo   iRepository.IUserRepository
	hasher hasher.IHasher
}

func NewCreateUserSvc(user apptypes.SignUpUser, repo iRepository.IUserRepository, hasher hasher.IHasher, logger *log.Logger) *CreateUserSvc {
	return &CreateUserSvc{user: user, repo: repo, hasher: hasher}
}

func (u *CreateUserSvc) Run() error {
	var err error

	validations := []func() error{
		u.checkName,
		u.checkSurname,
		u.checkEmail,
		u.checkSex,
		u.checkPassword,
	}

	for _, validate := range validations {
		if err = validate(); err != nil {
			return err
		}
	}

	var userEntryDate, userBournDate time.Time

	if userEntryDate, err = u.checkEntryDate(); err != nil {
		return err
	}

	if userBournDate, err = u.checkBournDate(); err != nil {
		return err
	}

	var hash_password []byte

	if hash_password, err = u.hasher.Hash([]byte(u.user.Password)); err != nil {
		return err
	}

	insert_User := models.NewUser(0, u.user.Name, u.user.Surname, u.user.Email, u.user.Professor, userEntryDate, userBournDate, "", string(u.user.Sex), string(hash_password), 1)

	return u.repo.Insert(insert_User)
}

func (u *CreateUserSvc) checkName() error {
	return datacheck.CheckName(u.user.Name)
}

func (u *CreateUserSvc) checkSurname() error {
	return datacheck.CheckSurname(u.user.Surname)
}

func (u *CreateUserSvc) checkEmail() error {
	return datacheck.CheckEmail(u.user.Email)
}

func (u *CreateUserSvc) checkEntryDate() (time.Time, error) {
	dateLayout := "2006-01-02 15:04:05"
	date, err := time.Parse(dateLayout, u.user.EntryDate)
	return date, err
}

func (u *CreateUserSvc) checkBournDate() (time.Time, error) {
	dateLayout := "2006-01-02"
	date, err := time.Parse(dateLayout, u.user.BournDate)
	return date, err
}

func (u *CreateUserSvc) checkSex() error {
	return datacheck.CheckSex(u.user.Sex)
}

func (u *CreateUserSvc) checkPassword() error {
	return datacheck.CheckPassword(u.user.Password)
}
