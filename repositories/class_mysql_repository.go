package repositories

import (
	"database/sql"
	"fmt"
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

func (*ClassMySQLRepository) GetClassByCreatorUserId(int) (*models.Class, error) {
	panic("unimplemented")
}

func (u *ClassMySQLRepository) GetClassById(id int) (*models.Class, error) {
	query := "SELECT  `id`, `title`, `description`, `creation_date`, `creator_user_id` FROM classes WHERE id= ?"
	row, err := u.db.Query(query, id)

	defer row.Close()

	if err != nil {
		return &models.Class{}, err
	}
	Class := &models.Class{}
	dateLayout := "2006-01-02 15:04:05"
	var creationDate string
	if row.Next() {

		if err := row.Scan(&Class.Id, &Class.Title, &Class.Description, &creationDate, &Class.UserCreatorId); err != nil {
			return &models.Class{}, err
		}
		Class.CreationDate, err = time.Parse(dateLayout, creationDate)
		if err != nil {
			return &models.Class{}, err
		}
	}

	if err = row.Err(); err != nil {
		return Class, err
	}
	query = "SELECT c.user_id FROM `class_has_user` c WHERE c.id_Class=  ?"
	rows1, err := u.db.Query(query, id)

	defer rows1.Close()

	if err != nil {
		return &models.Class{}, err
	}

	var studentUserId int

	for rows1.Next() {
		if err = rows1.Scan(&studentUserId); err != nil {
			return Class, err
		}
		Class.StudentUsers = append(Class.StudentUsers, studentUserId)
	}

	if err = rows1.Err(); err != nil {
		return Class, err
	}

	query = "SELECT ucec.editor_user_id FROM `user_can_edit_class` ucec WHERE ucec.class_id=  ?"
	rows2, err := u.db.Query(query, id)

	if err != nil {
		return &models.Class{}, err
	}

	var idEditorUserId int
	for rows2.Next() {
		if err := rows2.Scan(&idEditorUserId); err != nil {
			return Class, err
		}
		Class.EditorUsers = append(Class.EditorUsers, idEditorUserId)
	}
	return Class, nil
}

func (u *ClassMySQLRepository) Insert(Class *models.Class) error {
	statement := "INSERT INTO `classes` ( `title`, `description`, `creation_date`, `creator_user_id`) VALUES ( ?, ?, ?, ?)"
	_, err := u.db.Exec(statement, Class.Title, Class.Description, Class.CreationDate.String()[:19], Class.UserCreatorId)
	return err
}

func (u *ClassMySQLRepository) AddStudentUser(Class *models.Class, usuario *models.User) (*models.Class, error) {
	statement := "INSERT INTO `class_has_user`(`entry_date`, `user_id`, `class_id`) VALUES ( ?, ?, ?)"
	_, err := u.db.Exec(statement, time.Now().String()[:19], usuario.Id, Class.Id)
	if err != nil {
		return Class, err
	}
	Class.StudentUsers = append(Class.StudentUsers, usuario.Id)
	return Class, nil
}

func (u *ClassMySQLRepository) AddEditorUser(Class *models.Class, usuario *models.User) (*models.Class, error) {
	statement := "INSERT INTO `user_can_edit_class`(`entry_date`, `editor_user_id`, `class_id`) VALUES ( ?, ?, ?)"
	_, err := u.db.Exec(statement, time.Now().String()[:19], usuario.Id, Class.Id)
	if err != nil {
		return Class, err
	}
	Class.EditorUsers = append(Class.EditorUsers, usuario.Id)
	return Class, nil
}

func (u *ClassMySQLRepository) AddContent(class *models.Class, content *models.Content, position int) (*models.ClassHasContent, error) {
	statement := "INSERT INTO `class_has_content`(, `class_id`, `content_id`, `position`) VALUES (?, ?, ?)"
	result, err := u.db.Exec(statement, class.Id, content.Id, position)

	if err != nil {
		return nil, err
	}

	fmt.Println(result.LastInsertId())
	return nil, nil
}
