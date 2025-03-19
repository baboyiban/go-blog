package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB { // Make sure this is exported (capitalized)
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err.Error())
	}
	return db
}
