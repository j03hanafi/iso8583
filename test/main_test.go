package main

import (
	"bytes"
	"github.com/j03hanafi/iso8583/pkg/db"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Init(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	db.Server().ServeHTTP(response, request)
	expectedResponse := `{"message":"Hello World!","status":200}`
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.Code, "Invalid response code")
	assert.Equal(t, expectedResponse, string(bytes.TrimSpace(responseBody)))
	t.Log(response.Body.String())
}
