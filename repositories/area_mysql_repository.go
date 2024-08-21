package repositories

import (
	"database/sql"
	"time"

	"github.com/milnner/b_modules/errors"
	"github.com/milnner/b_modules/models"
)

type AreaMySQLRepository struct {
	db *sql.DB
}

func NewAreaMySQLRepository(db *sql.DB) (*AreaMySQLRepository, error) {
	if db == nil {
		return nil, errors.NewDatabaseConnectionError()
	}
	return &AreaMySQLRepository{db: db}, nil
}

func (u *AreaMySQLRepository) GetAreaById(area *models.Area) (err error) {
	var (
		row              *sql.Rows
		creationDatetime string
	)
	statement := "SELECT `title`, `description`, `owner_id`, `creation_datetime`, `activated` FROM `area` WHERE `id`=?"

	if row, err = u.db.Query(statement, area.Id); err != nil {
		return err
	}

	if row.Next() {
		if err := row.Scan(&area.Title, &area.Description, &area.OwnerId, &creationDatetime, &area.Activated); err != nil {
			return err
		}
		if area.CreationDatetime, err = time.Parse(time.DateTime, creationDatetime); err != nil {
			return err
		}
	}
	return err
}

func (u *AreaMySQLRepository) GetAreasByOwnerId(areas []models.Area, user *models.User) (err error) {
	var (
		row              *sql.Rows
		creationDatetime string
	)
	statement := "SELECT `title`, `description`, `owner_id`, `creation_datetime`, `activated` FROM `area` WHERE `owner_id`=?"

	if row, err = u.db.Query(statement, user.Id); err != nil {
		return err
	}

	for row.Next() {
		var area models.Area

		if err := row.Scan(&area.Title, &area.Description, &area.OwnerId, &creationDatetime, &area.Activated); err != nil {
			return err
		}
		if area.CreationDatetime, err = time.Parse(time.DateTime, creationDatetime); err != nil {
			return err
		}
		areas = append(areas, area)
	}
	return err
}

func (u *AreaMySQLRepository) GetUserIdsByAreaId(area *models.Area) (areaIds []int, err error) {
	var (
		row *sql.Rows
		id  int
	)
	statement := "SELECT `user_id` FROM `user_has_area_access` WHERE `area_id`=?"

	if row, err = u.db.Query(statement, area.Id); err != nil {
		return nil, err
	}

	for row.Next() {
		if err := row.Scan(&id); err != nil {
			return nil, err
		}
		areaIds = append(areaIds, id)
	}
	return areaIds, err
}

func (u *AreaMySQLRepository) InsertUser(area *models.Area, user *models.User) (err error) {
	statement := "INSERT INTO `user_has_area_access`(`permission`, `area_id`, `user_id`) VALUES (?,?,?)"
	_, err = u.db.Exec(statement, user.Permision, area.Id, user.Id)
	return err
}

func (u *AreaMySQLRepository) RemoveUser(area *models.Area, user *models.User) (err error) {
	statement := "UPDATE `user_has_area_access` SET `activated`=0 WHERE ( `area_id`, `user_id`)=(?,?)"
	_, err = u.db.Exec(statement, area.Id, user.Id)
	return err
}

func (u *AreaMySQLRepository) Insert(area *models.Area) (err error) {
	statement := "INSERT INTO `area`(`title`, `description`, `owner_id`) VALUES (?,?,?)"
	_, err = u.db.Exec(statement, area.Title, area.Description, area.OwnerId)
	return err
}
func (u *AreaMySQLRepository) Update(area *models.Area) (err error) {
	statement := "UPDATE `area` SET `title`=?,`description`=?, `owner_id`=? WHERE `id`=? and `owner_id`=?"
	_, err = u.db.Exec(statement, area.Title, area.Description, area.OwnerId, area.Id, area.OwnerId)
	return err
}
func (u *AreaMySQLRepository) Delete(area *models.Area) (err error) {
	statement := "UPDATE `area` SET `activated`=0 WHERE `id`=? and `owner_id`=?"
	_, err = u.db.Exec(statement, area.Id, area.OwnerId)
	return err
}

func (u *AreaMySQLRepository) GetPermission(area *models.Area, user *models.User) (err error) {
	var row *sql.Rows
	statement := "SELECT `permission` FROM `user_has_area_access` WHERE (`area_id`, `user_id`)=(?,?) and `activated`=1 "
	if row, err = u.db.Query(statement, area.Id, user.Id); err != nil {
		return err
	}
	if row.Next() {
		err = row.Scan(&user.Permision)
	}
	return err
}

func (u *AreaMySQLRepository) GetAreasByIds(areas []models.Area) (err error) {
	for i := 0; i < len(areas); i++ {
		if err = u.GetAreaById(&areas[i]); err != nil {
			return err
		}
	}
	return nil
}

func (u *AreaMySQLRepository) GetAreaIdsByOwnerId(area *models.Area) (areaIds []int, err error) {
	var (
		row *sql.Rows
		id  int
	)
	statement := "SELECT `id` FROM `area` WHERE `owner_id`=?"

	if row, err = u.db.Query(statement, area.OwnerId); err != nil {
		return nil, err
	}

	for row.Next() {
		if err := row.Scan(&id); err != nil {
			return nil, err
		}
		areaIds = append(areaIds, id)
	}
	return areaIds, err
}
