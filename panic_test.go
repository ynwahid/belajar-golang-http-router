package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func Test_Panic(t *testing.T) {
	router := httprouter.New()

	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, err interface{}) {
		fmt.Fprint(writer, "Panic: ", err)
	}

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("Ups...")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Panic: Ups...", string(body))
}
