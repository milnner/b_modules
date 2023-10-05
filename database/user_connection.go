package database

import (
	"database/sql"
	"fmt"
)

func GetDBUserInsert(sqlDriver string) (*sql.DB, error) {
	if insertUser == nil {
		if err := initDBConnection(&insertUser, connectionStringForInsertUser, sqlDriver); err != nil {
			return nil, err
		}
	}
	fmt.Println(insertContent)
	return insertUser, nil
}

func GetDBUserSelect(sqlDriver string) (*sql.DB, error) {
	if selectUser == nil {
		if err := initDBConnection(&selectUser, connectionStringForSelectUser, sqlDriver); err != nil {
			fmt.Println("Entrou")
			return nil, err
		}
	}
	return selectUser, nil
}
