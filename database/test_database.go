package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	testConnectionStringForTestInsertUser = "root:root@tcp(127.0.0.1:3306)/ardeo_test"
	testConnectionStringForSelectUser     = "root:root@tcp(127.0.0.1:3306)/ardeo_test"
	testConnectionStringForSelectClass    = "root:root@tcp(127.0.0.1:3306)/ardeo_test"
)

var (
	testInsertUser  *sql.DB
	testSelectUser  *sql.DB
	testSelectClass *sql.DB
)
