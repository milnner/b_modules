package repositories

import (
	"database/sql"
	"time"

	errapp "github.com/milnner/b_modules/errors"
	"github.com/milnner/b_modules/models"
)

type ContentMySQLRepository struct {
	db *sql.DB
}

func NewContentMySQLRepository(db *sql.DB) (*ContentMySQLRepository, error) {
	if db == nil {
		return nil, errapp.NewDatabaseConnectionError()
	}
	return &ContentMySQLRepository{db: db}, nil
}

func (u *ContentMySQLRepository) Insert(content *models.Content) (err error) {
	statement := "INSERT INTO `contents`( `title`, `description`, `area_id`) VALUES (?,?,?)"
	if _, err = u.db.Exec(statement, content.Title, content.Description, content.AreaId); err != nil {
		return err
	}
	return nil
}

func (u *ContentMySQLRepository) Update(content *models.Content) (err error) {
	statement := "UPDATE `contents` SET `title`=?,`description`=?,`last_update`=? WHERE `id`= ?"
	if _, err = u.db.Exec(statement, content.Title, content.Description, content.LastUpdate.String()[:19], content.Id); err != nil {
		return err
	}
	return nil
}

func (u *ContentMySQLRepository) Delete(content *models.Content) (err error) {
	statement := "UPDATE `contents` SET `activated`=0 WHERE `id`= ?"
	if _, err = u.db.Exec(statement, content.Id); err != nil {
		return err
	}
	return nil
}

func (u *ContentMySQLRepository) GetContentById(content *models.Content) (err error) {
	var row *sql.Rows
	query := "SELECT `id`, `title`, `description`, `creation_datetime`, `last_update`, `area_id`, `activated` FROM `contents` WHERE `id`=?"
	if row, err = u.db.Query(query, content.Id); err != nil {
		return err
	}

	if row.Next() {
		var (
			creationDatetime string
			lastUpdate       string
		)
		if err = row.Scan(&content.Id, &content.Title, &content.Description, &creationDatetime, &lastUpdate, &content.AreaId, &content.Activated); err != nil {
			return err
		}
		if content.CreationDate, err = time.Parse(time.DateTime, creationDatetime); err != nil {
			return err
		}
		if content.LastUpdate, err = time.Parse(time.DateTime, lastUpdate); err != nil {
			return err
		}
	}
	return nil
}

func (u *ContentMySQLRepository) GetContentsByAreaId(area *models.Area) (contents []models.Content, err error) {
	var (
		rows    *sql.Rows
		content models.Content
	)
	query := "SELECT `id`, `title`, `description`, `creation_datetime`, `last_update`, `area_id`, `activated` FROM `contents` WHERE `area_id`=?"

	if rows, err = u.db.Query(query, area.Id); err != nil {
		return nil, err
	}
	var (
		creationDatetime string
		lastUpdate       string
	)
	for rows.Next() {
		if err = rows.Scan(&content.Id, &content.Title, &content.Description, &creationDatetime, &lastUpdate, &content.AreaId, &content.Activated); err != nil {
			return nil, err
		}
		if content.CreationDate, err = time.Parse(time.DateTime, creationDatetime); err != nil {
			return nil, err
		}
		if content.LastUpdate, err = time.Parse(time.DateTime, lastUpdate); err != nil {
			return nil, err
		}
		contents = append(contents, content)
	}
	return contents, nil
}

func (u *ContentMySQLRepository) GetContentsByIds(contents []models.Content) (err error) {
	for i := 0; i < len(contents); i++ {
		if err = u.GetContentById(&contents[i]); err != nil {
			return err
		}
	}
	return nil
}

func (u *ContentMySQLRepository) AddActivity(content *models.Content, activity interface{}) (err error) {
	var (
		statement string
		args      []any
	)
	args = append(args, content.AreaId)
	args = append(args, content.Id)
	switch act := activity.(type) {
	case *models.TextActivity:
		statement = "INSERT INTO `content_see_activity`(`area_id`, `content_id`, `text_activity_id`, `position`) VALUES (?,?,?,?)"
		args = append(args, act.Id)
		args = append(args, act.Position)
	case *models.ImageActivity:
		statement = "INSERT INTO `content_see_activity`(`area_id`, `content_id`, `image_activity_id`, `position`) VALUES (?,?,?,?)"
		args = append(args, act.Id)
		args = append(args, act.Position)
	case *models.OneQuestionNAnswerActivity:
		statement = "INSERT INTO `content_see_activity`(`area_id`, `content_id`, `one_question_n_answer_activity_id`, `position`) VALUES (?,?,?,?)"
		args = append(args, act.Id)
		args = append(args, act.Position)

	}
	if _, err = u.db.Exec(statement, args...); err != nil {
		return err
	}
	return nil
}

func (u *ContentMySQLRepository) RemoveActivity(content *models.Content, activity interface{}) (err error) {
	var (
		statement string
		args      []any
	)
	args = append(args, content.AreaId)
	args = append(args, content.Id)

	switch act := activity.(type) {
	case *models.TextActivity:
		statement = "UPDATE `content_see_activity` SET `activated`=0 WHERE (`area_id`, `content_id`, `text_activity_id`)=(?,?,?)"
		args = append(args, act.Id)
	case *models.ImageActivity:
		statement = "UPDATE `content_see_activity` SET `activated`=0 WHERE (`area_id`, `content_id`, `image_activity_id`)=(?,?,?)"
		args = append(args, act.Id)
	case *models.OneQuestionNAnswerActivity:
		statement = "UPDATE `content_see_activity` SET `activated`=0 WHERE (`area_id`, `content_id`, `one_question_n_answer_activity_id`)=(?,?,?)"
		args = append(args, act.Id)
	}
	if _, err = u.db.Exec(statement, args...); err != nil {
		return err
	}
	return nil
}
func (u *ContentMySQLRepository) UpdateActivityPosition(content *models.Content, activity interface{}) (err error) {
	var (
		statement string
		args      []any
	)

	switch act := activity.(type) {
	case *models.TextActivity:
		statement = "UPDATE `content_see_activity` SET `position`=? WHERE (`text_activity_id`,`area_id`, `content_id`)=(?,?,?)"
		args = append(args, act.Position)
		args = append(args, act.Id)
	case *models.ImageActivity:
		statement = "UPDATE `content_see_activity` SET `position`=? WHERE (`image_activity_id`,`area_id`, `content_id`)=(?,?,?)"
		args = append(args, act.Position)
		args = append(args, act.Id)
	case *models.OneQuestionNAnswerActivity:
		statement = "UPDATE `content_see_activity` SET `position`=? WHERE (`one_question_n_answer_activity_id`,`area_id`, `content_id`)=(?,?,?)"
		args = append(args, act.Position)
		args = append(args, act.Id)
	}
	args = append(args, content.AreaId)
	args = append(args, content.Id)
	if _, err = u.db.Exec(statement, args...); err != nil {
		return err
	}
	return nil
}
func (u *ContentMySQLRepository) GetActivityIdsByContentId(content *models.Content, activity interface{}) (actIds []int, err error) {
	var (
		query string
		rows  *sql.Rows
		id    int
		ids   []int
	)
	switch activity.(type) {
	case *models.TextActivity:
		query = "SELECT `text_activity_id` FROM `content_see_activity` WHERE (`area_id`, `content_id`) = (?, ?)  and `text_activity_id` != \"NULL\""
	case *models.ImageActivity:
		query = "SELECT `image_activity_id` FROM `content_see_activity` WHERE (`area_id`, `content_id`) = (?, ?) and `image_activity_id`!=\"NULL\""
	case *models.OneQuestionNAnswerActivity:
		query = "SELECT `one_question_n_answer_activity_id` FROM `content_see_activity` WHERE (`area_id`, `content_id`) = (?, ?)  and `one_question_n_answer_activity_id`!= \"NULL\""
	}

	if rows, err = u.db.Query(query, content.AreaId, content.Id); err != nil {
		return nil, err
	}

	for rows.Next() {

		if err = rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}
