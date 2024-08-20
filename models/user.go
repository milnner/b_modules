package models

import (
	"reflect"
	"sort"
	"time"

	apptypes "github.com/milnner/b_modules/apptypes"
)

type Users []User

type User struct {
	Id        int                          `json:"id"`
	Name      string                       `json:"name"`
	Surname   string                       `json:"surname"`
	Email     string                       `json:"email"`
	Professor uint8                        `json:"professor"`
	EntryDate time.Time                    `json:"entryDate"`
	BournDate time.Time                    `json:"bournDate"`
	Sex       apptypes.Sex                 `json:"sex"`
	Permision apptypes.UserClassPermission `json:"permission"`
	Hash      string                       `json:"hash"`
	Activated uint8                        `json:"activated"`
}

func NewUser(id int, name string, surname string, email string, professor uint8, entryDate time.Time, bournDate time.Time, permision apptypes.UserClassPermission, sex string, hash string, activated uint8) *User {
	return &User{Id: id, Name: name, Surname: surname, Email: email, Professor: professor, EntryDate: entryDate, BournDate: bournDate, Sex: apptypes.Sex(sex), Hash: hash, Activated: activated, Permision: permision}

}

func (u User) Equals(other User) bool {
	return u.Id == other.Id &&
		u.Name == other.Name &&
		u.Surname == other.Surname &&
		u.Email == other.Email &&
		u.Professor == other.Professor &&
		u.EntryDate.Equal(other.EntryDate) &&
		u.BournDate.Equal(other.BournDate) &&
		u.Sex == other.Sex &&
		reflect.DeepEqual(u.Permision, other.Permision) &&
		u.Hash == other.Hash &&
		u.Activated == other.Activated
}

func (u Users) Sort(fieldName string) {
	sortUserByField(u, fieldName)
}

func sortUserByField(slice []User, fieldName string) {
	sort.Slice(slice, func(i, j int) bool {
		valI := reflect.ValueOf(slice[i])
		valJ := reflect.ValueOf(slice[j])

		fieldI := valI.FieldByName(fieldName)
		fieldJ := valJ.FieldByName(fieldName)

		return fieldI.Int() < fieldJ.Int()
	})
}
