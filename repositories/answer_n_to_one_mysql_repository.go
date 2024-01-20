package repositories

import (
	"database/sql"

	errapp "github.com/milnner/b_modules/errors"
	"github.com/milnner/b_modules/models"
)

type AnswerNToOneMySQLRepository struct {
	db *sql.DB
}

func NewAnswerNToOneMySQLRepository(db *sql.DB) (*AnswerNToOneMySQLRepository, error) {
	if db == nil {
		return nil, errapp.NewDatabaseConnectionError()
	}
	return &AnswerNToOneMySQLRepository{db: db}, nil
}

func (u *AnswerNToOneMySQLRepository) Insert(answerNToOne *models.AnswerNToOne) (err error) {
	statement := "INSERT INTO `answer_n_to_one`(`area_id`, `one_question_n_answer_activity_id`, `correctness`, `answer` ) VALUES (?,?,?,?)"
	if _, err = u.db.Exec(statement, answerNToOne.AreaId, answerNToOne.OneQuestionNAnswerActivityId, answerNToOne.Correctness, answerNToOne.Answer); err != nil {
		return err
	}
	return nil
}

func (u *AnswerNToOneMySQLRepository) Update(answerNToOne *models.AnswerNToOne) (err error) {
	statement := "UPDATE `answer_n_to_one` SET `correctness`=?,`answer`=? WHERE `id`=?"
	if _, err = u.db.Exec(statement, answerNToOne.Correctness, answerNToOne.Answer, answerNToOne.Id); err != nil {
		return err
	}
	return nil
}

func (u *AnswerNToOneMySQLRepository) Delete(answerNToOne *models.AnswerNToOne) (err error) {
	statement := "UPDATE `answer_n_to_one` SET `activated`=? WHERE `id`=?"
	if _, err = u.db.Exec(statement, 0, answerNToOne.Id); err != nil {
		return err
	}
	return nil
}

func (u *AnswerNToOneMySQLRepository) GetAnswerNToOneById(answerNToOne *models.AnswerNToOne) (err error) {
	var row *sql.Rows
	query := "SELECT `area_id`, `one_question_n_answer_activity_id`, `correctness`, `answer`, `activated` FROM `answer_n_to_one` WHERE `id`=?"

	if row, err = u.db.Query(query, answerNToOne.Id); err != nil {
		return err
	}
	if row.Next() {
		if err = row.Scan(&answerNToOne.AreaId, &answerNToOne.OneQuestionNAnswerActivityId, &answerNToOne.Correctness, &answerNToOne.Answer, &answerNToOne.Activated); err != nil {
			return err
		}
	}
	return nil
}

func (u *AnswerNToOneMySQLRepository) GetAnswersNToOneByIds(answersNToOne []models.AnswerNToOne) (err error) {
	for i := 0; i < len(answersNToOne); i++ {
		if err = u.GetAnswerNToOneById(&answersNToOne[i]); err != nil {
			return err
		}
	}
	return nil
}

func (u *AnswerNToOneMySQLRepository) GetAnswersNToOneByOneQuestionNAnswerActivityId(oneQuestionNAnswerActivity *models.OneQuestionNAnswerActivity) (answersNToOne []models.AnswerNToOne, err error) {
	var (
		answerNToOne models.AnswerNToOne
		rows         *sql.Rows
	)
	query := "SELECT `id`, `area_id`, `one_question_n_answer_activity_id`, `correctness`, `answer`, `activated` FROM `answer_n_to_one` WHERE `one_question_n_answer_activity_id`=?"

	if rows, err = u.db.Query(query, oneQuestionNAnswerActivity.Id); err != nil {
		return nil, err
	}
	for rows.Next() {
		if err = rows.Scan(&answerNToOne.Id, &answerNToOne.AreaId, &answerNToOne.OneQuestionNAnswerActivityId, &answerNToOne.Correctness, &answerNToOne.Answer, &answerNToOne.Activated); err != nil {
			return nil, err
		}

		answersNToOne = append(answersNToOne, answerNToOne)
	}
	return answersNToOne, nil
}

func (u *AnswerNToOneMySQLRepository) GetAnswersNToOneIdsByOneQuestionNAnswerActivityId(oneQuestionNAnswerActivity *models.OneQuestionNAnswerActivity) (answersIds []int, err error) {
	var (
		id   int
		rows *sql.Rows
	)
	query := "SELECT `id` FROM `answer_n_to_one` WHERE `one_question_n_answer_activity_id`=?"

	if rows, err = u.db.Query(query, oneQuestionNAnswerActivity.Id); err != nil {
		return nil, err
	}
	for rows.Next() {
		if err = rows.Scan(&id); err != nil {
			return nil, err
		}

		answersIds = append(answersIds, id)
	}
	return answersIds, err
}
