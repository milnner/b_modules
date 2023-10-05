package models

import (
	"time"

	apptypes "github.com/milnner/b_modules/apptypes"
)

type User struct {
	Id        int
	Name      string
	Surname   string
	Email     string
	EntryDate time.Time
	BournDate time.Time
	Sex       apptypes.Sex
	Hash      string
}

func NewUser(name string, surname string, email string, entry_date time.Time, bourn_date time.Time, sex string, hash string) *User {
	return &User{Name: name, Surname: surname, Email: email, EntryDate: entry_date, BournDate: bourn_date, Sex: apptypes.Sex(sex), Hash: hash}

}
