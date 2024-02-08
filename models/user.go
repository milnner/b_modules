package models

import (
	"reflect"
	"sort"
	"time"

	apptypes "github.com/milnner/b_modules/apptypes"
)

type Users []User

type User struct {
	Id        int
	Name      string
	Surname   string
	Email     string
	Professor uint8
	EntryDate time.Time
	BournDate time.Time
	Sex       apptypes.Sex
	Permision apptypes.UserClassPermission
	Hash      string
	Activated uint8
}

func NewUser(id int, name string, surname string, email string, professor uint8, entryDate time.Time, bournDate time.Time, permision apptypes.UserClassPermission, sex string, hash string, activated uint8) *User {
	return &User{Id: id, Name: name, Surname: surname, Email: email, Professor: professor, EntryDate: entryDate, BournDate: bournDate, Sex: apptypes.Sex(sex), Hash: hash, Activated: activated, Permision: permision}

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
