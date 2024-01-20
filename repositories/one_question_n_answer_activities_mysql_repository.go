package repositories

import (
	"database/sql"
	"time"

	errapp "github.com/milnner/b_modules/errors"
	"github.com/milnner/b_modules/models"
)

type OneQuestionNAnswerActivityMySQLRepository struct {
	db *sql.DB
}

func NewOneQuestionNAnswerActivityMySQLRepository(db *sql.DB) (*OneQuestionNAnswerActivityMySQLRepository, error) {
	if db == nil {
		return nil, errapp.NewDatabaseConnectionError()
	}
	return &OneQuestionNAnswerActivityMySQLRepository{db: db}, nil
}

func (u *OneQuestionNAnswerActivityMySQLRepository) Insert(oneQuestionNAnswerActivity *models.OneQuestionNAnswerActivity) (err error) {
	statement := "INSERT INTO `one_question_n_answer_activities`(`area_id`, `question`) VALUES (?,?)"
	if _, err = u.db.Exec(statement, oneQuestionNAnswerActivity.AreaId, oneQuestionNAnswerActivity.Question); err != nil {
		return err
	}
	return nil
}

func (u *OneQuestionNAnswerActivityMySQLRepository) GetOneQuestionNAnswerActivityById(oneQuestionNAnswerActivity *models.OneQuestionNAnswerActivity) (err error) {
	var (
		rows       *sql.Rows
		lastUpdate string
	)
	query := "SELECT `id`, `area_id`, `question`, `last_update`, `activated` FROM `one_question_n_answer_activities` WHERE `id`=?"
	if rows, err = u.db.Query(query, oneQuestionNAnswerActivity.Id); err != nil {
		return err
	}
	if rows.Next() {
		if err = rows.Scan(&oneQuestionNAnswerActivity.Id, &oneQuestionNAnswerActivity.AreaId, &oneQuestionNAnswerActivity.Question, &lastUpdate, &oneQuestionNAnswerActivity.Activated); err != nil {
			return err
		}
		if oneQuestionNAnswerActivity.LastUpdate, err = time.Parse(time.DateTime, lastUpdate); err != nil {
			return err
		}
	}
	return nil
}
func (u *OneQuestionNAnswerActivityMySQLRepository) GetOneQuestionNAnswerActivitiesByIds(oneQuestionNAnswerActivity []models.OneQuestionNAnswerActivity) (err error) {
	for i := 0; i < len(oneQuestionNAnswerActivity); i++ {
		if err = u.GetOneQuestionNAnswerActivityById(&oneQuestionNAnswerActivity[i]); err != nil {
			return err
		}
	}
	return nil
}
func (u *OneQuestionNAnswerActivityMySQLRepository) GetOneQuestionNAnswerActivitiesByAreaId(area *models.Area) (oneQuestionNAnswerActivities []models.OneQuestionNAnswerActivity, err error) {
	var (
		rows                       *sql.Rows
		lastUpdate                 string
		oneQuestionNAnswerActivity models.OneQuestionNAnswerActivity
	)
	query := "SELECT `id`, `area_id`, `question`, `last_update`, `activated` FROM `one_question_n_answer_activities` WHERE `area_id`=? ORDER BY `id`"

	if rows, err = u.db.Query(query, area.Id); err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&oneQuestionNAnswerActivity.Id, &oneQuestionNAnswerActivity.AreaId, &oneQuestionNAnswerActivity.Question, &lastUpdate, &oneQuestionNAnswerActivity.Activated); err != nil {
			return nil, err
		}
		if oneQuestionNAnswerActivity.LastUpdate, err = time.Parse(time.DateTime, lastUpdate); err != nil {
			return nil, err
		}
		oneQuestionNAnswerActivities = append(oneQuestionNAnswerActivities, oneQuestionNAnswerActivity)
	}
	return oneQuestionNAnswerActivities, nil
}

func (u *OneQuestionNAnswerActivityMySQLRepository) Delete(oneQuestionNAnswerActivity *models.OneQuestionNAnswerActivity) (err error) {
	statement := "UPDATE `one_question_n_answer_activities` SET `activated`=? WHERE `id`=?"

	if _, err = u.db.Exec(statement, 0, oneQuestionNAnswerActivity.Id); err != nil {
		return err
	}
	return nil
}
func (u *OneQuestionNAnswerActivityMySQLRepository) Update(oneQuestionNAnswerActivity *models.OneQuestionNAnswerActivity) (err error) {
	statement := "UPDATE `one_question_n_answer_activities` SET `question`=?,`last_update`=?,`activated`=? WHERE `id`=?"

	if _, err = u.db.Exec(statement, oneQuestionNAnswerActivity.Question, oneQuestionNAnswerActivity.LastUpdate.String()[:19], oneQuestionNAnswerActivity.Activated, oneQuestionNAnswerActivity.Id); err != nil {
		return err
	}
	return nil
}

func (u *OneQuestionNAnswerActivityMySQLRepository) GetOneQuestionNAnswerActivityIdsByAreaId(area *models.Area) (oneQstNAswIds []int, err error) {
	var (
		rows *sql.Rows
		id   int
	)
	query := "SELECT `id` FROM `one_question_n_answer_activities` WHERE `area_id`=? "

	if rows, err = u.db.Query(query, area.Id); err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&id); err != nil {
			return nil, err
		}

		oneQstNAswIds = append(oneQstNAswIds, id)
	}
	return oneQstNAswIds, err
}
