package repositories

import (
	"database/sql"
	"time"

	errapp "github.com/milnner/b_modules/errors"
	models "github.com/milnner/b_modules/models"
)

type userMySQLRepository struct {
	db *sql.DB
}

func NewUserMySQLRepository(db *sql.DB) (*userMySQLRepository, error) {
	if db == nil {
		return nil, errapp.NewDatabaseConnectionError()
	}
	return &userMySQLRepository{db: db}, nil
}

func (u *userMySQLRepository) GetUserById(id int) (*models.User, error) {
	query := "SELECT id, name, surname, email, entry_date, bourn_date, sex, hash FROM users WHERE id= ?"
	row, err := u.db.Query(query, id)

	if err != nil {
		return &models.User{}, err
	}

	user := &models.User{}

	var entry_date, bourn_date string
	if row.Next() {
		if err = row.Scan(&user.Id, &user.Name, &user.Surname, &user.Email, &entry_date, &bourn_date, &user.Sex, &user.Hash); err != nil {
			return &models.User{}, err
		}
	}
	dateLayout := "2006-01-02 15:04:05"

	user.EntryDate, err = time.Parse(dateLayout, entry_date)

	if err != nil {
		return &models.User{}, err
	}
	user.BournDate, err = time.Parse(dateLayout, bourn_date)
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}

func (u *userMySQLRepository) GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT id, name, surname, email, entry_date, bourn_date, sex, hash FROM users WHERE email= ?"
	row, err := u.db.Query(query, email)

	if err != nil {
		return &models.User{}, err
	}
	user := &models.User{}

	var entry_date, bourn_date string
	if row.Next() {

		if err = row.Scan(&user.Id, &user.Name, &user.Surname, &user.Email, &entry_date, &bourn_date, &user.Sex, &user.Hash); err != nil {
			return &models.User{}, err
		}
	}
	dateLayout := "2006-01-02 15:04:05"

	user.BournDate, err = time.Parse(dateLayout, bourn_date)

	if err != nil {
		return &models.User{}, err
	}

	user.EntryDate, err = time.Parse(dateLayout, entry_date)

	if err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func (u *userMySQLRepository) GetAll() (*[]models.User, error) {
	query := "SELECT id, name, surname, email, entry_date, bourn_date, sex, hash FROM users"
	rows, err := u.db.Query(query)

	if err != nil {
		return nil, err
	}
	user := models.User{}
	users := []models.User{}

	var entry_date, bourn_date string
	if rows.Next() {

		if err = rows.Scan(&user.Id, &user.Name, &user.Surname, &user.Email, &entry_date, &bourn_date, &user.Sex, &user.Hash); err != nil {
			return nil, err
		}

		dateLayout := "2006-01-02 15:04:05"

		user.BournDate, err = time.Parse(dateLayout, bourn_date)

		if err != nil {
			return nil, err
		}

		user.EntryDate, err = time.Parse(dateLayout, entry_date)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return &users, nil
}

func (u *userMySQLRepository) Insert(user *models.User) error {
	statement := "INSERT INTO `users` ( `name`, `surname`, `email`, `entry_date`, `bourn_date`, `sex`, `hash`) VALUES ( ?, ?, ?, ?, ?, ?, ?)"

	_, err := u.db.Exec(statement, user.Name, user.Surname, user.Email, user.EntryDate.String()[:19], user.BournDate.String()[:10], user.Sex, user.Hash)
	return err
}

func (u *userMySQLRepository) Delete(user *models.User) error {
	statement := " DELETE FROM `users` WHERE id= ?"
	_, err := u.db.Exec(statement, user.Id)
	if err != nil {
		return err
	}
	return nil
}
