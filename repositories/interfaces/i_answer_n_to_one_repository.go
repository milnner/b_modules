package interfaces

import "github.com/milnner/b_modules/models"

type IAnswerNToOneRepository interface {
	GetAnswerNToOneById(*models.AnswerNToOne) error
	GetAnswersNToOneByIds([]models.AnswerNToOne) error
	GetAnswersNToOneByOneQuestionNAnswerActivityId(*models.OneQuestionNAnswerActivity) ([]models.AnswerNToOne, error)
	Insert(*models.AnswerNToOne) error
	Delete(*models.AnswerNToOne) error
	Update(*models.AnswerNToOne) error
}
