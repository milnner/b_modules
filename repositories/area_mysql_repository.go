package repositories

import (
	"database/sql"
	"time"

	"github.com/milnner/b_modules/errors"
	models "github.com/milnner/b_modules/models"
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

// Implementação da interface IAreaRepository
func (r *AreaMySQLRepository) GetAll() ([]models.Area, error) {
	query := "SELECT `id`, `title`, `description`, `owner_id`, `creation_datetime` FROM `area`"
	areas, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	var (
		area             models.Area
		creationDatetime string
		areasSlice       []models.Area
	)

	for areas.Next() {
		if err := areas.Scan(&area.Id, &area.Title, &area.Description, &area.OwnerId, &creationDatetime); err != nil {
			return nil, err
		}
		area.CreationDatetime, err = time.Parse(time.DateTime, creationDatetime)
		if err != nil {
			return nil, err
		}
		areasSlice = append(areasSlice, area)
	}
	return areasSlice, nil
}

func (r *AreaMySQLRepository) GetAreaById(area *models.Area) (err error) {
	query := "SELECT `id`, `title`, `description`, `owner_id`, `creation_datetime` FROM `area` WHERE `id`=?"
	var areas *sql.Rows
	areas, err = r.db.Query(query, area.Id)
	if err != nil {
		return err
	}
	var (
		creationDatetime string
	)

	if areas.Next() {
		if err := areas.Scan(&area.Id, &area.Title, &area.Description, &area.OwnerId, &creationDatetime); err != nil {
			return err
		}
		area.CreationDatetime, err = time.Parse(time.DateTime, creationDatetime)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *AreaMySQLRepository) GetAreasByUserId(user *models.User) ([]models.Area, error) {
	query := "SELECT `id`, `title`, `description`, `owner_id`, `creation_datetime` FROM `area` WHERE `owner_id`=?"
	areas, err := r.db.Query(query, user.Id)
	if err != nil {
		return nil, err
	}
	var (
		area             models.Area
		creationDatetime string
		areasArr         []models.Area
	)

	for areas.Next() {
		if err := areas.Scan(&area.Id, &area.Title, &area.Description, &area.OwnerId, &creationDatetime); err != nil {
			return nil, err
		}
		area.CreationDatetime, err = time.Parse(time.DateTime, creationDatetime)
		if err != nil {
			return nil, err
		}
		areasArr = append(areasArr, area)
	}
	return areasArr, nil
}

func (r *AreaMySQLRepository) GetAreaClassIdsById(class *models.Class) (ids []int, err error) {
	query := "SELECT `id` FROM `classes` WHERE `area_id`=?"
	rows, err := r.db.Query(query, class.Id)
	if err != nil {
		return nil, err
	}
	var idS int
	for rows.Next() {
		if err := rows.Scan(&idS); err != nil {
			return nil, err
		}
		ids = append(ids, idS)
	}
	return ids, nil
}

func (r *AreaMySQLRepository) GetAreaContentIdsById(content *models.Content) (ids []int, err error) {
	query := "SELECT `id` FROM `content` WHERE `area_id`=?"
	rows, err := r.db.Query(query, content.Id)
	if err != nil {
		return nil, err
	}
	var idS int
	for rows.Next() {
		if err := rows.Scan(&idS); err != nil {
			return nil, err
		}
		ids = append(ids, idS)
	}
	return ids, nil
}

func (r *AreaMySQLRepository) Insert(area *models.Area) error {
	query := "INSERT INTO `area`(`title`, `description`, `owner_id`, `creation_datetime`) VALUES (?,?,?,?)"
	_, err := r.db.Query(query, area.Title, area.Description, area.OwnerId, area.CreationDatetime)
	if err != nil {
		return err
	}
	area = nil
	return nil
}

func (r *AreaMySQLRepository) Update(area *models.Area) error {
	query := "UPDATE `area` SET `title`=?, `description`=?,`owner_id`=?,`creation_datetime`=? WHERE `id`=?"
	_, err := r.db.Query(query, area.Title, area.Description, area.OwnerId, area.CreationDatetime, area.Id)
	if err != nil {
		return err
	}
	area = nil
	return nil
}

func (r *AreaMySQLRepository) GetAreasByIds(ids []int) (areas []models.Area, err error) {
	var area models.Area
	for _, v := range ids {
		area.Id = v
		err = r.GetAreaById(&area)
		if err != nil {
			return areas, err
		}
		areas = append(areas, area)
	}
	return areas, nil
}

func (r *AreaMySQLRepository) Delete(area *models.Area) error {
	query := "UPDATE SET `activated`=0 WHERE id= ?"
	_, err := r.db.Query(query, area.Id)
	if err != nil {
		return err
	}
	area = nil
	return nil
}
