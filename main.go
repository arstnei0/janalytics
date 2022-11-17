package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env.local")
}

func main() {
	initDb()
	createTables()
	pages = make(map[string]Page)
	writeNumberString := os.Getenv("DB_WRITE_NUMBER")
	var writeNumberProcessing int
	if writeNumberString == "" {
		writeNumberProcessing = 5
	} else {
		writeNumberProcessing, _ = strconv.Atoi(writeNumberString)
	}

	writeNumber = uint32(writeNumberProcessing) / 2

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "GET"},
	}))

	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.POST("/site", func(ctx *gin.Context) {
		b, jsonErr, ioErr := readBody[Site](ctx)

		if ifJSONErrRespondErrElse(jsonErr, ctx) && ioErr == nil {
			_, dbErr := db.Exec("INSERT INTO Site VALUES ($1, $2)", b.Id, b.Name)

			if ifDbErrRespondErrElse(dbErr, ctx) {
				ctx.String(200, "OK")
			}
		}
	})

	r.GET("/sites", func(ctx *gin.Context) {
		rows := failIfFuncErr(db.Query("SELECT id, name FROM Site"))

		defer rows.Close()

		sites := []Site{}

		for rows.Next() {
			row := Site{}
			rows.Scan(&row.Id, &row.Name)
			sites = append(sites, row)
		}

		ctx.JSON(200, sites)
	})

	r.GET("/:site/:id", func(ctx *gin.Context) {
		siteId := ctx.Param("site")
		pageId := ctx.Param("id")

		page := viewPage(ctx, siteId, pageId)

		// fmt.Println(page.Views)

		ctx.JSON(200, page.Views)
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "hello")
	})

	r.Run(fmt.Sprintf(":%s", port))
}
