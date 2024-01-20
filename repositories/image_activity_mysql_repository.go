package repositories

import (
	"database/sql"
	"encoding/hex"
	"time"

	errapp "github.com/milnner/b_modules/errors"
	"github.com/milnner/b_modules/models"
)

type ImageActivityMySQLRepository struct {
	db *sql.DB
}

func NewImageActivityMySQLRepository(db *sql.DB) (*ImageActivityMySQLRepository, error) {
	if db == nil {
		return nil, errapp.NewDatabaseConnectionError()
	}
	return &ImageActivityMySQLRepository{db: db}, nil
}

func (u *ImageActivityMySQLRepository) Insert(imageActivity *models.ImageActivity) (err error) {
	query := "INSERT INTO `image_activities`(`area_id`, `title`, `_blob`) VALUES (?, ?, ?)"

	if _, err = u.db.Exec(query, imageActivity.AreaId, imageActivity.Title, hex.EncodeToString(imageActivity.Blob)); err != nil {
		return err
	}
	return nil
}

func (u *ImageActivityMySQLRepository) Delete(imageActivity *models.ImageActivity) (err error) {
	query := "UPDATE `image_activities` SET `activated`=? WHERE `id`=?"
	_, err = u.db.Exec(query, 0, imageActivity.Id)
	if err != nil {
		return err
	}
	return nil
}

func (u *ImageActivityMySQLRepository) Update(imageActivity *models.ImageActivity) (err error) {
	query := "UPDATE `image_activities` SET `title`=?,`_blob`=?,`last_update`=? WHERE `id`=?"
	_, err = u.db.Exec(query, imageActivity.Title, imageActivity.Blob, imageActivity.LastUpdate.String()[:19], imageActivity.Id)
	if err != nil {
		return err
	}
	return nil
}

func (u *ImageActivityMySQLRepository) GetImageActivityById(imageActivity *models.ImageActivity) (err error) {
	var (
		rows       *sql.Rows
		lastUpdate string
	)
	query := "SELECT  `area_id`, `title`, `_blob`, `last_update`, `activated` FROM `image_activities` WHERE id=?"
	rows, err = u.db.Query(query, imageActivity.Id)

	if err != nil {
		return err
	}

	if rows.Next() {
		if err = rows.Scan(&imageActivity.AreaId, &imageActivity.Title, &imageActivity.Blob, &lastUpdate, &imageActivity.Activated); err != nil {
			return err
		}
		if imageActivity.LastUpdate, err = time.Parse(time.DateTime, lastUpdate); err != nil {
			return err
		}
	}
	return nil
}

func (u *ImageActivityMySQLRepository) GetImageActivitiesByIds(imageActivities []models.ImageActivity) (err error) {

	for i := 0; i < len(imageActivities); i++ {
		if err = u.GetImageActivityById(&imageActivities[i]); err != nil {
			return err
		}
	}
	return nil
}

func (u *ImageActivityMySQLRepository) GetImageActivitiesByAreaId(area *models.Area) (imageActivities []models.ImageActivity, err error) {
	var (
		imageActivity models.ImageActivity
		rows          *sql.Rows
		lastUpdate    string
	)

	query := "SELECT `id`, `area_id`, `title`, `_blob`, `last_update`, `activated` FROM `image_activities` WHERE area_id=?"

	if rows, err = u.db.Query(query, area.Id); err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&imageActivity.Id, &imageActivity.AreaId, &imageActivity.Title, &imageActivity.Blob, &lastUpdate, &imageActivity.Activated); err != nil {
			return nil, err
		}
		if imageActivity.LastUpdate, err = time.Parse(time.DateTime, lastUpdate); err != nil {
			return nil, err
		}
		imageActivities = append(imageActivities, imageActivity)
	}
	return imageActivities, nil
}
func (u *ImageActivityMySQLRepository) GetImageActivityIdsByAreaId(area *models.Area) (imgIds []int, err error) {
	var (
		rows *sql.Rows
		id   int
	)
	query := "SELECT `id` FROM `image_activities` WHERE `area_id`=?"

	if rows, err = u.db.Query(query, area.Id); err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&id); err != nil {
			return nil, err
		}
		imgIds = append(imgIds, id)
	}
	return imgIds, err
}
