package database

import (
	"database/sql"
)

func GetDBClassSelect(sqlDriver string) (*sql.DB, error) {
	if selectClass == nil {
		if err := initDBConnection(&selectClass, connectionStringForSelectClass, sqlDriver); err != nil {
			return nil, err
		}
	}
	return selectClass, nil
}

func GetDBClassInsert(sqlDriver string) (*sql.DB, error) {
	if insertClass == nil {
		if err := initDBConnection(&insertClass, connectionStringForInsertClass, sqlDriver); err != nil {
			return nil, err
		}
	}
	return insertClass, nil
}
