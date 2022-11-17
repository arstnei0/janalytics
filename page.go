package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Site  string
	Id    string `json:"id"`
	Views uint32 `json:"views"`
}

func createPageTable() {
	failIfFuncErr(db.Exec("DROP TABLE IF EXISTS Page"))

	failIfFuncErr(db.Exec(`CREATE TABLE Page (
		id text,
		site VARCHAR(50),
		views serial
	)`))

	dbLog("Table Site Created.")
}

var pages map[string]Page
var writeNumber uint32

func viewPage(ctx *gin.Context, siteId string, pageId string) Page {
	if page, ok := pages[pageId]; ok {
		page.Views += 1
		pages[pageId] = page
	} else {
		page, err := getPage(siteId, pageId)

		if err == sql.ErrNoRows {
			page, _ := createPage(siteId, pageId)
			page.Views += 1
			pages[pageId] = page
		} else {
			page.Views += 1
			pages[pageId] = page
		}
	}

	page := pages[pageId]

	if page.Views%writeNumber == 0 {
		go func() {
			_, err := db.Exec("UPDATE page SET views=views+$1 WHERE id=$2 AND site=$3", writeNumber, pageId, siteId)
			failIfErr(err)
		}()
	}

	return page
}

func createPage(siteId string, pageId string) (Page, error) {
	page := Page{
		Id:    pageId,
		Site:  siteId,
		Views: 0,
	}

	_, err := db.Exec("INSERT INTO Page VALUES ($1, $2, $3)", page.Id, page.Site, page.Views)

	return page, err
}

func getPage(siteId string, pageId string) (Page, error) {
	result := db.QueryRow("SELECT * FROM Page WHERE site=$1 AND id=$2", siteId, pageId)
	var site string
	var id string
	var views uint32
	err := result.Scan(&site, &id, &views)
	page := Page{
		Site:  site,
		Id:    id,
		Views: views,
	}

	return page, err
}
