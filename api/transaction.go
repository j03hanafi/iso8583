package api

import (
	"github.com/gorilla/mux"
	"github.com/j03hanafi/iso8583/pkg/db"
	"github.com/j03hanafi/iso8583/pkg/helper"
	"net/http"
)

func GetPayments(writer http.ResponseWriter, request *http.Request) {
	var response db.PaymentsResponse
	err := db.PingDb(db.MysqlDB)

	if err != nil {
		response.ResponseStatus.ReasonCode, response.ResponseStatus.ResponseDescription = 500, db.DatabaseError
		helper.JsonFormatter(writer, response, 500)
	}

	payments, statusCode, err := db.SelectPayments(response, db.MysqlDB)
	response.TransactionData = payments
	response.ResponseStatus.ReasonCode, response.ResponseStatus.ResponseDescription = statusCode, "Success"

	helper.JsonFormatter(writer, response, statusCode)

}

func GetPayment(writer http.ResponseWriter, request *http.Request) {
	var response db.PaymentResponse
	err := db.PingDb(db.MysqlDB)
	processingCode := mux.Vars(request)["processingCode"]

	if err != nil {
		response.ResponseStatus.ReasonCode, response.ResponseStatus.ResponseDescription = 500, db.DatabaseError
		helper.JsonFormatter(writer, response, 500)
	}

	payment, statusCode, err := db.SelectPayment(response, processingCode, db.MysqlDB)
	response.TransactionData = payment
	response.ResponseStatus.ReasonCode, response.ResponseStatus.ResponseDescription = statusCode, "Success"

	helper.JsonFormatter(writer, response, statusCode)
}

func CreatePayment(writer http.ResponseWriter, request *http.Request) {

}

func UpdatePayment(writer http.ResponseWriter, request *http.Request) {

}

func DeletePayment(writer http.ResponseWriter, request *http.Request) {

}
