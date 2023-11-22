package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// ex: err := initDBConnection(conn, "user:password@tcp(host:port)/database", "driver");
func InitDBConnection(conn **sql.DB, stringForConn, sqlDriverName string) error {
	var err error
	*conn, err = sql.Open(sqlDriverName, stringForConn)
	if err != nil {
		return err
	}
	if err = (*conn).Ping(); err != nil {
		return err
	}
	return nil
}
