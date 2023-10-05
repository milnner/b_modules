package database

import (
	"database/sql"
)

func GetDBClassSelectTest(sqlDriver string) (*sql.DB, error) {
	if testSelectClass == nil {
		if err := initDBConnection(&testSelectClass, testConnectionStringForSelectClass, sqlDriver); err != nil {
			return nil, err
		}
	}
	return testSelectClass, nil
}
