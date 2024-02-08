package repositories

import (
	"database/sql"
	"time"

	errapp "github.com/milnner/b_modules/errors"
	models "github.com/milnner/b_modules/models"
)

type UserMySQLRepository struct {
	db *sql.DB
}

func NewUserMySQLRepository(db *sql.DB) (*UserMySQLRepository, error) {
	if db == nil {
		return nil, errapp.NewDatabaseConnectionError()
	}
	return &UserMySQLRepository{db: db}, nil
}

func (u *UserMySQLRepository) GetUserById(user *models.User) (err error) {
	var (
		row *sql.Rows
	)

	query := "SELECT `id`, `name`, `surname`, `email`, `professor`, `entry_date`, `bourn_date`, `sex`, `hash`, `activated` FROM `users` WHERE `id`=?"

	if row, err = u.db.Query(query, user.Id); err != nil {
		return err
	}

	var (
		entryDate string
		bournDate string
	)
	if row.Next() {
		if err = row.Scan(&user.Id,
			&user.Name,
			&user.Surname,
			&user.Email,
			&user.Professor,
			&entryDate,
			&bournDate,
			&user.Sex,
			&user.Hash,
			&user.Activated); err != nil {
			return err
		}
		if user.EntryDate, err = time.Parse(time.DateTime, entryDate); err != nil {
			return err
		}

		if user.BournDate, err = time.Parse(time.DateTime, bournDate); err != nil {
			return err
		}

	}

	return err
}

func (u *UserMySQLRepository) GetUserByEmail(user *models.User) (err error) {
	var (
		row *sql.Rows
	)

	query := "SELECT `id`, `name`, `surname`, `email`, `professor`, `entry_date`, `bourn_date`, `sex`, `hash`, `activated` FROM `users` WHERE `email`= ?"

	if row, err = u.db.Query(query, user.Email); err != nil {
		return err
	}

	var (
		entryDate string
		bournDate string
	)
	if row.Next() {
		if err = row.Scan(&user.Id,
			&user.Name,
			&user.Surname,
			&user.Email,
			&user.Professor,
			&entryDate,
			&bournDate,
			&user.Sex,
			&user.Hash,
			&user.Activated); err != nil {
			return err
		}
		if user.EntryDate, err = time.Parse(time.DateTime, entryDate); err != nil {
			return err
		}

		if user.BournDate, err = time.Parse(time.DateTime, bournDate); err != nil {
			return err
		}

	}

	return err
}

func (u *UserMySQLRepository) Update(user *models.User) (err error) {
	statement := "UPDATE `users` SET `name`=?,`surname`=?,`email`=?, `professor`=?, `entry_date`=?,`bourn_date`=?,`sex`=?,`hash`=?,`activated`=? WHERE `id`=?"
	_, err = u.db.Exec(statement, user.Name, user.Surname, user.Email, user.Professor, user.EntryDate.String()[:19], user.BournDate.String()[:10], user.Sex, user.Hash, user.Activated, user.Id)
	return err
}

func (u *UserMySQLRepository) Insert(user *models.User) error {
	statement := "INSERT INTO `users` ( `name`, `surname`, `email`, `professor`, `entry_date`, `bourn_date`, `sex`, `hash`) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := u.db.Exec(statement, user.Name, user.Surname, user.Email, user.Professor, user.EntryDate.String()[:19], user.BournDate.String()[:10], user.Sex, user.Hash)
	return err
}

func (u *UserMySQLRepository) Delete(user *models.User) error {
	statement := "UPDATE `users` SET `activated`=0 WHERE id= ?"
	_, err := u.db.Exec(statement, user.Id)

	return err
}

func (u *UserMySQLRepository) GetUsersByIds(users []models.User) (err error) {
	for i := 0; i < len(users); i++ {
		if err = u.GetUserById(&users[i]); err != nil {
			return err
		}
	}
	return err
}
