package models

type AnswerNToOne struct {
	Id                           int
	AreaId                       int
	OneQuestionNAnswerActivityId int
	Correctness                  uint8
	Answer                       []byte
	Activated                    uint8
}

func NewAnswerNToOne(id, areaId, oneQuestionNAnswerActivityId int, correctness uint8, answer []byte, activated uint8) *AnswerNToOne {
	return &AnswerNToOne{Id: id, AreaId: areaId, OneQuestionNAnswerActivityId: oneQuestionNAnswerActivityId, Correctness: correctness, Answer: answer, Activated: activated}
}
