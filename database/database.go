package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// Class database connection strings
	connectionStringForSelectClass = "root:root@tcp(127.0.0.1:3306)/ardeo"
	connectionStringForInsertClass = "root:root@tcp(127.0.0.1:3306)/ardeo"

	// User database connection strings
	connectionStringForInsertUser = "root:root@tcp(127.0.0.1:3306)/ardeo"
	connectionStringForSelectUser = "root:root@tcp(127.0.0.1:3306)/ardeo"

	// Content database connection strings
	connectionStringForInsertContent = "root:root@tcp(127.0.0.1:3306)/ardeo"
)

var (
	// User database connetion
	insertUser *sql.DB
	selectUser *sql.DB

	// Class database connection
	insertClass *sql.DB
	selectClass *sql.DB

	// Content database connection
	insertContent *sql.DB
)

// ex: err := initDBConnection(conn, "user:password@tcp(host:port)/database", "driver");
func initDBConnection(conn **sql.DB, stringForConn, sqlDriverName string) error {
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
