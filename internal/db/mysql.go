package db

import (
	"database/sql"
	"log"
	"os"
)

func CreateMySqlDB() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DB"))
	if err != nil {
		log.Fatalln(err.Error())
	}

	return db
}
