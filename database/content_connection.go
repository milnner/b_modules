package database

import "database/sql"

func GetDBContentInsert(sqlDriver string) (*sql.DB, error) {
	if insertContent == nil {
		if err := initDBConnection(&insertContent, connectionStringForInsertContent, sqlDriver); err != nil {
			return nil, err
		}
	}
	return insertContent, nil
}
