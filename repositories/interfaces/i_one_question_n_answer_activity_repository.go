package interfaces

import "github.com/milnner/b_modules/models"

type IOneQuestionNAnswerActivityRepository interface {
	GetOneQuestionNAnswerActivityById(*models.OneQuestionNAnswerActivity) error
	GetOneQuestionNAnswerActivitiesByIds([]models.OneQuestionNAnswerActivity) error
	GetOneQuestionNAnswerActivitiesByAreaId(*models.Area) ([]models.OneQuestionNAnswerActivity, error)
	GetOneQuestionNAnswerActivityIdsByAreaId(*models.Area) ([]int, error)
	Insert(*models.OneQuestionNAnswerActivity) error
	Delete(*models.OneQuestionNAnswerActivity) error
	Update(*models.OneQuestionNAnswerActivity) error
}
