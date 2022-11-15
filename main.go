package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type Domain struct {
	views map[string]int
}

func main() {
	// domains := make(map[string]Domain)

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

	r.POST("/domain", func(c *gin.Context) {
		fmt.Println("aaa")

		fmt.Println(c.Request.Body)
		c.JSON(200, "aaa")
	})

	r.Run(fmt.Sprintf(":%s", port))
}
