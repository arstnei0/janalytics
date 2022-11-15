package main

import (
	"encoding/json"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

// Check If Error
func cie(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func cfie[T any](result T, err error) T {
	cie(err)

	return result
}

func readBody[T any](c *gin.Context) (T, error, error) {
	var result T
	body, ioErr := io.ReadAll(c.Request.Body)

	jsonErr := json.Unmarshal(body, &result)

	return result, ioErr, jsonErr
}

func ifDbErrRespondErr(err error, ctx *gin.Context) {
	if err != nil {
		ctx.String(500, "DB Error!")
	}
}

func ifErrRespondErr(err error, ctx *gin.Context) {
	if err != nil {
		ctx.String(500, "DB Error!")
	}
}
