package models

import "time"

type OneQuestionNAnswerActivity struct {
	Id         int
	AreaId     int
	Question   []byte
	LastUpdate time.Time
	Activated  uint8
	Position   int
}

func NewOneQuestionNAnswerActivity(id, areaId int, question []byte, lastUpdate time.Time, activated uint8) *OneQuestionNAnswerActivity {
	return &OneQuestionNAnswerActivity{Id: id, AreaId: areaId, Question: question, LastUpdate: lastUpdate, Activated: activated}
}
