package api

import (
	"github.com/j03hanafi/iso8583/pkg/db"
	"github.com/j03hanafi/iso8583/pkg/helper"
	"net/http"
)

func InitHandler(writer http.ResponseWriter, request *http.Request) {
	var response db.Transaction
	helper.JsonFormatter(writer, map[string]interface{}{
		"status":      200,
		"message":     "Hello World!",
		"transaction": response,
	}, 200)
}
