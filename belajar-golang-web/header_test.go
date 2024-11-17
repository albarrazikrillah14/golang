package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")

	fmt.Fprint(w, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "Medomeckz")

	fmt.Println(w, "Ok")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	_, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	header := recorder.Header()

	fmt.Println(header)
}
