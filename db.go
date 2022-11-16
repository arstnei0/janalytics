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

	db = failIfFuncErr(sql.Open("postgres", psqlInfo))
}

func createTables() {
	if os.Getenv("DB_RESET") == "true" {
		createSiteTable()
		createPageTable()
	}
}
