package main

type Site struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func createSiteTable() {
	failIfFuncErr(db.Exec("DROP TABLE IF EXISTS Site"))

	failIfFuncErr(db.Exec(`CREATE TABLE Site (
		id VARCHAR(50) PRIMARY KEY,
		name text
	)`))

	dbLog("Table Page Created.")
}
