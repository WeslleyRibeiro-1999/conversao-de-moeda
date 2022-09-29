package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:weslley123@tcp(localhost:3306)/conversao")
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	return db, err
}
