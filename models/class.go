package models

import "time"

type Class struct {
	Id            int
	Title         string
	Description   string
	CreationDate  time.Time
	UserCreatorId int
	AreaId        int
	LastUpdate    time.Time
	Activated     uint8
}

func NewClass(id int, title string, description string, creationDate time.Time, userCreatorId int, areaId int, lastUpdate time.Time, activated uint8) *Class {
	return &Class{Id: id, Title: title, Description: description, CreationDate: creationDate, UserCreatorId: userCreatorId, AreaId: areaId, LastUpdate: lastUpdate, Activated: activated}
}
