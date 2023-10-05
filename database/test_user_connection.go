package database

import (
	"database/sql"
)

func GetDBUserInsertTest(sqlDriver string) (*sql.DB, error) {
	if testInsertUser == nil {
		if err := initDBConnection(&testInsertUser, testConnectionStringForTestInsertUser, sqlDriver); err != nil {

			return nil, err
		}
	}
	return testInsertUser, nil
}

func GetDBUserSelectTest(sqlDriver string) (*sql.DB, error) {
	if testSelectUser == nil {
		if err := initDBConnection(&testSelectUser, testConnectionStringForSelectUser, sqlDriver); err != nil {
			return nil, err
		}
	}
	return testSelectUser, nil
}
