package models

import "time"

type Content struct {
	Id           int
	CreationDate time.Time
	Title        string
	Description  string
	LastUpdate   time.Time
	AreaId       int
	Activated    uint8
	Position     uint8
}

func NewContent(id int, creationDate time.Time, title, description string, lastUpdate time.Time, areaId int, activated uint8) *Content {
	return &Content{Id: id, CreationDate: creationDate, Title: title, Description: description, LastUpdate: lastUpdate, AreaId: areaId, Activated: activated}
}
