package main

import (
	"fmt"
	"os"

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

	// r.GET("/view/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	v, prs := views[id]
	// 	if !prs {
	// 		views[id] = 0
	// 	}

	// 	newV := v + 1
	// 	views[id] = newV

	// 	c.JSON(200, newV)
	// })

	r.POST("/site", func(ctx *gin.Context) {
		b, ioErr, jsonErr := readBody[Site](ctx)
		_, dbErr := db.Exec("INSERT INTO Site VALUES ($1, $2)", b.Id, b.Name)

		if ioErr != nil {
			ctx.String(400, "IO Error")
		}

		if jsonErr != nil {
			ctx.String(400, "JSON Format not right")
		}

		ctx.JSON(200, "OK")
	})

	r.GET("/sites", func(ctx *gin.Context) {
		rows := cfie(db.Query("SELECT id, name FROM Site"))

		defer rows.Close()

		sites := []Site{}

		for rows.Next() {
			row := Site{}
			rows.Scan(&row.Id, &row.Name)
			sites = append(sites, row)
		}

		ctx.JSON(200, sites)
	})

	r.Run(fmt.Sprintf(":%s", port))
}
