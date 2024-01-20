package repositories

import (
	"database/sql"
	"time"

	errapp "github.com/milnner/b_modules/errors"
	models "github.com/milnner/b_modules/models"
)

type ClassMySQLRepository struct {
	db *sql.DB
}

func NewClassMySQLRepository(db *sql.DB) (*ClassMySQLRepository, error) {
	if db == nil {
		return nil, errapp.NewDatabaseConnectionError()
	}
	return &ClassMySQLRepository{db: db}, nil
}

func (u *ClassMySQLRepository) GetClassById(class *models.Class) (err error) {
	var row *sql.Rows
	query := "SELECT `id`, `title`, `description`, `creation_datetime`, `creator_user_id`, `last_update`, `area_id`, `activated` FROM `classes` WHERE id= ?"

	defer func() {
		err = row.Close()
	}()

	if row, err = u.db.Query(query, class.Id); err != nil {
		return err
	}
	var (
		creationDate string
		lastUpdate   string
	)

	if row.Next() {
		if err := row.Scan(&class.Id, &class.Title, &class.Description, &creationDate, &class.UserCreatorId, &lastUpdate, &class.AreaId, &class.Activated); err != nil {
			return err
		}
		if class.CreationDate, err = time.Parse(time.DateTime, creationDate); err != nil {
			return err
		}
		if class.LastUpdate, err = time.Parse(time.DateTime, lastUpdate); err != nil {
			return err
		}
	}
	return err
}

func (u *ClassMySQLRepository) Update(class *models.Class) (err error) {
	statement := "UPDATE `classes` SET  `title`=?, `description`=?, `last_update`=? WHERE `id`=?"
	_, err = u.db.Exec(statement, class.Title, class.Description, time.Now().String()[:19], class.Id)
	return err
}

func (u *ClassMySQLRepository) Delete(class *models.Class) (err error) {
	statement := "UPDATE `classes` SET  `activated`=0 WHERE `id`=?"
	_, err = u.db.Exec(statement, class.Id)
	return err
}

func (u *ClassMySQLRepository) Insert(class *models.Class) error {
	statement := "INSERT INTO `classes` ( `title`, `description`, `creator_user_id`, `area_id`) VALUES ( ?, ?, ?, ?)"
	_, err := u.db.Exec(statement, class.Title, class.Description, class.UserCreatorId, class.AreaId)
	return err
}

func (u *ClassMySQLRepository) AddContent(class *models.Class, content *models.Content) (err error) {
	statement := "INSERT INTO `class_see_content`(`class_id`, `content_id`, `position`) VALUES (?,?,?)"
	_, err = u.db.Exec(statement, class.Id, content.Id, content.Position)
	return err
}

func (u *ClassMySQLRepository) RemoveContent(class *models.Class, content *models.Content) (err error) {
	statement := "UPDATE `class_see_content` SET `activated`=0 WHERE (`content_id`, `class_id`)=(?,?)"
	_, err = u.db.Exec(statement, class.Id, content.Id)
	return err
}

func (u *ClassMySQLRepository) UpdateContentPosition(class *models.Class, content *models.Content) (err error) {
	statement := "UPDATE `class_see_content` SET `position`=? WHERE (`content_id`, `class_id`)=(?,?)"
	_, err = u.db.Exec(statement, content.Position, class.Id, content.Id)
	return err
}

func (u *ClassMySQLRepository) GetContentIdsById(class *models.Class) (contentIds []int, err error) {
	var (
		rows *sql.Rows
		id   int
	)
	statement := "SELECT `content_id` FROM `class_see_content` WHERE `class_id`=?"
	if rows, err = u.db.Query(statement, class.Id); err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&id); err != nil {
			return nil, err
		}
		contentIds = append(contentIds, id)
	}

	return contentIds, err
}

func (u *ClassMySQLRepository) AddStudentUser(class *models.Class, user *models.User) (err error) {
	statement := "INSERT INTO `class_takes_user`(`user_id`, `class_id`) VALUES (?,?)"
	_, err = u.db.Exec(statement, user.Id, class.Id)
	return err
}

func (u *ClassMySQLRepository) RemoveStudentUser(class *models.Class, user *models.User) (err error) {
	statement := "UPDATE `class_takes_user` SET `activated`=0 WHERE (`user_id`,`class_id`)=(?,?)"
	_, err = u.db.Exec(statement, user.Id, class.Id)
	return err

}

func (u *ClassMySQLRepository) GetStudentIdsById(class *models.Class) (userIds []int, err error) {
	var (
		rows *sql.Rows
		id   int
	)
	statement := "SELECT `user_id`FROM `class_takes_user` WHERE `class_id`=?"

	if rows, err = u.db.Query(statement, class.Id); err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&id); err != nil {
			return nil, err
		}
		userIds = append(userIds, id)
	}

	return userIds, err
}

func (u *ClassMySQLRepository) GetClassesByIds(classes []models.Class) (err error) {
	for i := 0; i < len(classes); i++ {
		if err = u.GetClassById(&classes[i]); err != nil {
			return err
		}
	}
	return nil
}

func (u *ClassMySQLRepository) GetClassIdsByAreaId(area *models.Area) (classIds []int, err error) {
	var (
		id  int
		row *sql.Rows
	)
	query := "SELECT `id` FROM `classes` WHERE `area_id`= ?"

	defer func() {
		err = row.Close()
	}()

	if row, err = u.db.Query(query, area.Id); err != nil {
		return nil, err
	}

	for row.Next() {
		if err := row.Scan(&id); err != nil {
			return nil, err
		}
		classIds = append(classIds, id)
	}
	return classIds, err
}
