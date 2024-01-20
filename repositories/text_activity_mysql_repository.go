package repositories

import (
	"database/sql"
	"encoding/hex"
	"time"

	errapp "github.com/milnner/b_modules/errors"
	"github.com/milnner/b_modules/models"
)

type TextActivityMySQLRepository struct {
	db *sql.DB
}

func NewTextActivityMySQLRepository(db *sql.DB) (*TextActivityMySQLRepository, error) {
	if db == nil {
		return nil, errapp.NewDatabaseConnectionError()
	}
	return &TextActivityMySQLRepository{db: db}, nil
}

func (u *TextActivityMySQLRepository) Insert(textActivity *models.TextActivity) (err error) {
	statement := "INSERT INTO `text_activities`(`area_id`, `title`, `_blob`) VALUES (?, ?, ?)"

	if _, err = u.db.Exec(statement, textActivity.AreaId, textActivity.Title, hex.EncodeToString(textActivity.Blob)); err != nil {
		return err
	}
	return nil
}

func (u *TextActivityMySQLRepository) Delete(textActivity *models.TextActivity) (err error) {
	statement := "UPDATE `text_activities` SET `activated`=0 WHERE `id`=?"
	if _, err = u.db.Exec(statement, textActivity.Id); err != nil {
		return err
	}
	textActivity = nil
	return nil
}

func (u *TextActivityMySQLRepository) Update(textActivity *models.TextActivity) (err error) {
	statement := "UPDATE `text_activities` SET `title`=?,`_blob`=?,`last_update`=? WHERE `id`=?"
	if _, err = u.db.Exec(statement, textActivity.Title, textActivity.Blob, textActivity.LastUpdate.String()[:19], textActivity.Id); err != nil {
		return err
	}
	return nil
}

func (u *TextActivityMySQLRepository) GetTextActivityById(textActivity *models.TextActivity) (err error) {
	query := "SELECT `area_id`, `title`, `_blob`, `last_update`, `activated` FROM `text_activities` WHERE `id`=?"
	var textActivityRow *sql.Rows
	if textActivityRow, err = u.db.Query(query, textActivity.Id); err != nil {
		return err
	}

	var lastUpdateStr string
	if textActivityRow.Next() {
		if err = textActivityRow.Scan(&textActivity.AreaId, &textActivity.Title, &textActivity.Blob, &lastUpdateStr, &textActivity.Activated); err != nil {
			return err
		}
		var lastUpdate time.Time
		if lastUpdate, err = time.Parse(time.DateTime, lastUpdateStr); err != nil {
			return err
		}
		textActivity.LastUpdate = lastUpdate
	}
	return nil
}

func (u *TextActivityMySQLRepository) GetTextActivitiesByIds(textActivities []models.TextActivity) (err error) {
	for i := 0; i < len(textActivities); i++ {
		if err = u.GetTextActivityById(&textActivities[i]); err != nil {
			return err
		}
	}
	return nil
}

func (u *TextActivityMySQLRepository) GetTextActivitiesByAreaId(area *models.Area) (textActivities []models.TextActivity, err error) {
	var (
		textActivityRow *sql.Rows
		lastUpdateStr   string
		textActivity    models.TextActivity
	)
	query := "SELECT `id`, `area_id`, `title`, `_blob`, `last_update`, `activated` FROM `text_activities` WHERE area_id=?"

	if textActivityRow, err = u.db.Query(query, area.Id); err != nil {
		return nil, err
	}

	for textActivityRow.Next() {
		if err = textActivityRow.Scan(&textActivity.Id, &textActivity.AreaId, &textActivity.Title, &textActivity.Blob, &lastUpdateStr, &textActivity.Activated); err != nil {
			return nil, err
		}
		var lastUpdate time.Time
		if lastUpdate, err = time.Parse(time.DateTime, lastUpdateStr); err != nil {
			return nil, err
		}
		textActivity.LastUpdate = lastUpdate
		textActivities = append(textActivities, textActivity)
	}
	return textActivities, nil
}
func (u *TextActivityMySQLRepository) GetTextActivityIdsByAreaId(area *models.Area) (txtIds []int, err error) {
	var (
		textActivityRow *sql.Rows
		id              int
	)
	query := "SELECT `id` FROM `text_activities` WHERE `area_id`=?"

	if textActivityRow, err = u.db.Query(query, area.Id); err != nil {
		return nil, err
	}

	for textActivityRow.Next() {
		if err = textActivityRow.Scan(&id); err != nil {
			return nil, err
		}

		txtIds = append(txtIds, id)
	}
	return txtIds, err
}
