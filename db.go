package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func dbLog(a ...any) {
	fmt.Print("[DB] ")
	fmt.Println(a...)
}

func initDb() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db = cfie(sql.Open("postgres", psqlInfo))
}

func createTables() {
	cfie(db.Exec("DROP TABLE IF EXISTS Site"))

	cfie(db.Exec(`CREATE TABLE Site (
		id VARCHAR(50) PRIMARY KEY,
		name text
	)`))

	dbLog("Table Site Created.")
}
