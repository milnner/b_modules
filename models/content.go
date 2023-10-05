package models

import "time"

type Content struct {
	Id            int
	CreatorUserId int
	CreationDate  time.Time
	Title         string
	Description   string
}

func NewContent(id, creatorUserId int, creationDate time.Time, title, description string) *Content {
	return &Content{Id: id, CreatorUserId: creatorUserId, CreationDate: creationDate, Title: title, Description: description}
}
