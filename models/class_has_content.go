package models

type ClassHasContent struct {
	Id        int
	ClassId   int
	ContentId int
	Position  int
}

func NewClassHasContent(id, classId, contentId, position int) *ClassHasContent {
	return &ClassHasContent{Id: id, ClassId: classId, ContentId: contentId, Position: position}
}
