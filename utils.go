package main

import (
	"encoding/json"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

// Check If Error
func failIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func failIfFuncErr[T any](result T, err error) T {
	failIfErr(err)

	return result
}

func readBody[T any](ctx *gin.Context) (T, error, error) {
	var result T
	body, ioErr := io.ReadAll(ctx.Request.Body)

	jsonErr := json.Unmarshal(body, &result)

	if ifIOErrRespondErrElse(ioErr, ctx) {
		return result, jsonErr, nil
	}

	return result, jsonErr, ioErr
}

func ifDbErrRespondErrElse(err error, ctx *gin.Context) bool {
	if err != nil {
		ctx.String(500, "DB Error!")
		return false
	}

	return true
}

func ifJSONErrRespondErrElse(err error, ctx *gin.Context) bool {
	if err != nil {
		ctx.String(400, "JSON Format Wrong!")
		return false
	}

	return true
}

func ifIOErrRespondErrElse(err error, ctx *gin.Context) bool {
	if err != nil {
		ctx.String(400, "IO Error!")
		return false
	}

	return true
}

func ifNoErr(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}
