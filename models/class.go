package models

import "time"

type Class struct {
	Id            int
	Title         string
	Description   string
	CreationDate  time.Time
	UserCreatorId int
	EditorUsers   []int
	StudentUsers  []int
}

func NewClass(id int, title string, description string, creationDate time.Time, userCreatorId int, editorUsers []int, studentUsers []int) *Class {
	return &Class{Id: id, Title: title, Description: description, CreationDate: creationDate, UserCreatorId: userCreatorId, EditorUsers: editorUsers, StudentUsers: studentUsers}
}
