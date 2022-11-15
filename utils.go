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

func readBody[T any](ctx *gin.Context) (T, error) {
	var result T
	body, ioErr := io.ReadAll(ctx.Request.Body)

	jsonErr := json.Unmarshal(body, &result)

	if ioErr != nil {
		ctx.String(400, "IO Error!")
	}

	return result, jsonErr
}

func ifDbErrRespondErr(err error, ctx *gin.Context) {
	if err != nil {
		ctx.String(500, "DB Error!")
	}
}

func ifJSONErrRespondErr(err error, ctx *gin.Context) {
	if err != nil {
		ctx.String(400, "JSON Format Wrong!")
	}
}

func ifIOErrRespondErr(err error, ctx *gin.Context) {
	if err != nil {
		ctx.String(400, "IO Error!")
	}
}

func ifNoErr(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}
