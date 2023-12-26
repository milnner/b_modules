package models

import (
	"reflect"
	"sort"
	"time"
)

// encoding jpeg
type ImageActivity struct {
	Id         int
	AreaId     int
	Title      string
	Blob       []byte
	LastUpdate time.Time
	Activated  uint8
}

func NewImageActivity(id, areaId int, title string, blob []byte, last_update time.Time, activityId int) *ImageActivity {
	return &ImageActivity{Id: id, AreaId: areaId, Title: title, Blob: blob, LastUpdate: last_update, Activated: uint8(activityId)}
}

type ImageActivities []ImageActivity

func (u ImageActivities) Sort(fieldName string) {
	sortImageActivityByField(u, fieldName)
}

func sortImageActivityByField(slice []ImageActivity, fieldName string) {
	sort.Slice(slice, func(i, j int) bool {
		valI := reflect.ValueOf(slice[i])
		valJ := reflect.ValueOf(slice[j])

		fieldI := valI.FieldByName(fieldName)
		fieldJ := valJ.FieldByName(fieldName)

		return fieldI.Int() < fieldJ.Int()
	})
}
