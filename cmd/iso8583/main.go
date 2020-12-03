package main

import (
	"github.com/gorilla/mux"
	"github.com/j03hanafi/iso8583/api"
	"github.com/j03hanafi/iso8583/pkg/db"
	"github.com/j03hanafi/iso8583/pkg/iso"
	"log"
	"net/http"
)

func main() {
	db.MysqlDB = db.Connect()
	defer db.MysqlDB.Close()
	router := Server()
	log.Fatal(http.ListenAndServe(":8000", router))
}

func Server() *mux.Router {
	router := mux.NewRouter()

	//endpoints
	//CRUD
	router.HandleFunc("/", api.InitHandler)
	router.HandleFunc("/payment", api.GetPayments).Methods("GET")
	router.HandleFunc("/payment/{processingCode}", api.GetPayment).Methods("GET")
	router.HandleFunc("/payment/", api.CreatePayment).Methods("POST")
	router.HandleFunc("/payment/{processingCode}", api.UpdatePayment).Methods("PUT")
	router.HandleFunc("/payment/{processingCode}", api.DeletePayment).Methods("DELETE")

	//iso8583
	router.HandleFunc("/payment/{processingCode}/iso8583", iso.ToIso8583).Methods("GET")
	router.HandleFunc("/payment/{processingCode}/iso8583/{element}", iso.ExtractElem).Methods("GET")

	return router
}
