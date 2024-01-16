package models

import (
	"reflect"
	"sort"
	"time"
)

type Area struct {
	Id               int
	Title            string
	Description      string
	OwnerId          int
	CreationDatetime time.Time
	Activated        uint8
}

type Areas []Area

func (u Areas) Sort(fieldName string) {
	sortAreaByField(u, fieldName)
}

func NewArea(id int, title string, description string, ownerId int, creationDatetime time.Time) *Area {
	return &Area{Id: id, Title: title, Description: description, OwnerId: ownerId, CreationDatetime: creationDatetime}
}

func sortAreaByField(slice []Area, fieldName string) {
	sort.Slice(slice, func(i, j int) bool {
		valI := reflect.ValueOf(slice[i])
		valJ := reflect.ValueOf(slice[j])

		fieldI := valI.FieldByName(fieldName)
		fieldJ := valJ.FieldByName(fieldName)

		return fieldI.Int() < fieldJ.Int()
	})
}
