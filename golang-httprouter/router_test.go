package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello World")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	response := recoder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
