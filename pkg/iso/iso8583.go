package iso

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/j03hanafi/iso8583/pkg/db"
	"github.com/j03hanafi/iso8583/pkg/helper"
	"github.com/mofax/iso8583"
	"github.com/rivo/uniseg"
	"net/http"
	"strconv"
)

func convertIso(processingCode string) (iso8583.IsoStruct, int, error) {
	var payment db.PaymentResponse
	transaction, statusCode, err := db.SelectPayment(payment, processingCode, db.MysqlDB)
	cardAcceptor := transaction.CardAcceptorData.CardAcceptorName + transaction.CardAcceptorData.CardAcceptorCity + transaction.CardAcceptorData.CardAcceptorCountryCode

	one := iso8583.NewISOStruct("pkg/iso/iso_spec/spec1987.yml", false)

	if one.Mti.String() != "" {
		fmt.Printf("Empty generates invalid MTI")
	}

	one.AddMTI("0200")
	one.AddField(2, transaction.Pan)
	one.AddField(3, transaction.ProcessingCode)
	one.AddField(4, strconv.Itoa(transaction.TotalAmount))
	one.AddField(5, transaction.SettlementAmount)
	one.AddField(6, transaction.CardholderBillingAmount)
	one.AddField(7, transaction.TransmissionDateTime)
	one.AddField(9, transaction.SettlementConversionRate)
	one.AddField(10, transaction.CardHolderBillingConvRate)
	one.AddField(11, transaction.Stan)
	one.AddField(12, transaction.LocalTransactionTime)
	one.AddField(13, transaction.LocalTransactionDate)
	one.AddField(17, transaction.CaptureDate)
	one.AddField(18, transaction.CategoryCode)
	one.AddField(22, transaction.PointOfServiceEntryMode)
	one.AddField(37, transaction.Refnum)
	one.AddField(41, transaction.CardAcceptorData.CardAcceptorTerminalId)
	one.AddField(43, cardAcceptor)
	one.AddField(48, transaction.AdditionalData)
	one.AddField(49, transaction.Currency)
	one.AddField(50, transaction.SettlementCurrencyCode)
	one.AddField(51, transaction.CardHolderBillingCurrencyCode)
	one.AddField(57, transaction.AdditionalDataNational)

	return one, statusCode, err
}

func ToIso8583(writer http.ResponseWriter, request *http.Request) {
	var response db.Iso8583
	err := db.PingDb(db.MysqlDB)
	processingCode := mux.Vars(request)["processingCode"]

	if err != nil {
		response.ResponseStatus.ReasonCode, response.ResponseStatus.ResponseDescription = 500, db.DatabaseError
		helper.JsonFormatter(writer, response, 500)
	}

	isomsg, statusCode, _ := convertIso(processingCode)

	msg, _ := isomsg.ToString()

	mti := isomsg.Mti.String()
	//bitmap := one.Bitmap
	header := fmt.Sprintf("%04d", uniseg.GraphemeClusterCount(msg))

	response.Iso8583 = header + msg
	response.MTI = mti
	//response.Bitmap = bitmap
	response.ResponseStatus.ReasonCode, response.ResponseStatus.ResponseDescription = statusCode, "Success"

	helper.JsonFormatter(writer, response, statusCode)
}

func ExtractElem(writer http.ResponseWriter, request *http.Request) {
	var response db.ExtractElem
	err := db.PingDb(db.MysqlDB)
	processingCode := mux.Vars(request)["processingCode"]
	element := mux.Vars(request)["element"]
	elementInt, _ := strconv.Atoi(element)
	elementInt64 := int64(elementInt)

	if err != nil {
		response.ResponseStatus.ReasonCode, response.ResponseStatus.ResponseDescription = 500, db.DatabaseError
		helper.JsonFormatter(writer, response, 500)
	}

	isomsg, statusCode, _ := convertIso(processingCode)

	elementMap := isomsg.Elements.GetElements()
	data, exist := elementMap[elementInt64]

	var statusDesc string
	if exist {
		response.Data = data
		statusDesc = "Success"
	} else {
		response.Data = "Data not found"
		statusCode = 404
		statusDesc = "Data not found"
	}

	response.Element = elementInt64
	response.ResponseStatus.ReasonCode, response.ResponseStatus.ResponseDescription = statusCode, statusDesc

	helper.JsonFormatter(writer, response, statusCode)
}
