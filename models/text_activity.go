package models

import "time"

type TextActivity struct {
	Id         int
	AreaId     int
	Title      string
	Blob       []byte // encoding - utf-8
	LastUpdate time.Time
	Activated  uint8
	Position   int
}

func NewTextActivity(id int, areaId int, title string, blob []byte, lastUpdate time.Time, activated uint8) *TextActivity {
	return &TextActivity{Id: id, AreaId: areaId, Title: title, Blob: blob, LastUpdate: lastUpdate, Activated: activated}
}
