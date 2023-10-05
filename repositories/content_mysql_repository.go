package repositories

import (
	"database/sql"
	"time"

	errapp "github.com/milnner/b_modules/errors"
	models "github.com/milnner/b_modules/models"
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

func (u *ContentMySQLRepository) GetContentById(id int) (*models.Content, error) {
	query := "SELECT `id`, `user_creator_id`, `title`, `description`, `creation_date` FROM `content` WHERE `id`= ?"
	row, err := u.db.Query(query, id)
	if err != nil {
		return &models.Content{}, err
	}

	content := &models.Content{}

	if row.Next() {
		var creationDate string
		if err := row.Scan(&content.Id, &content.CreatorUserId, &content.Title, &content.Description, &creationDate); err != nil {
			return nil, err
		}
		content.CreationDate, err = time.Parse(time.DateTime, creationDate)
		if err != nil {
			return nil, err
		}
	}
	return content, nil
}

func (u *ContentMySQLRepository) GetContentByCreatorUserId(id int) (*models.Content, error) {
	query := "SELECT `id`, `user_creator_id`, `title`, `description`, `creation_date` FROM `content` WHERE `user_creator_id`= ?"
	row, err := u.db.Query(query, id)
	if err != nil {
		return &models.Content{}, err
	}

	content := &models.Content{}

	if row.Next() {
		var creationDate string

		if err := row.Scan(&content.Id, &content.CreatorUserId, &content.Title, &content.Description, &creationDate); err != nil {
			return nil, err
		}
		content.CreationDate, err = time.Parse(time.DateTime, creationDate)
		if err != nil {
			return nil, err
		}
	}
	return content, nil
}

func (u *ContentMySQLRepository) Insert(content *models.Content) error {
	statement := "INSERT INTO `content`(`id`, `user_creator_id`, `title`, `description`, `creation_date`) VALUES (?, ?, ?, ?, ?)"
	_, err := u.db.Exec(statement, content.Id, content.CreatorUserId, content.Title, content.Description, content.CreationDate)

	if err != nil {
		return err
	}
	return nil
}
