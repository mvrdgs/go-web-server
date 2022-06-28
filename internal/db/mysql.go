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

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS go_web_server;")
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = db.Exec("USE go_web_server;")
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS sellers (
    	id BINARY(16) PRIMARY KEY NOT NULL,
    	cid VARCHAR(60) NOT NULL UNIQUE,
		company_name VARCHAR(60) NOT NULL,
    	address VARCHAR(60) NOT NULL,
    	telephone VARCHAR(15) NOT NULL
	)`)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return db
}
