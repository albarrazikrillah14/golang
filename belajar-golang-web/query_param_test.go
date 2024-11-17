package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	fmt.Fprintf(w, "Hello %s", name)
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/helo?name=Zikri", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	bodyString := string(body)

	fmt.Println(bodyString)
}

func MultipleQueryParameter(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("firstName")
	lastName := r.URL.Query().Get("lastName")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?firstName=Albarra&lastName=Zikrillah", nil)

	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	bodyString := string(body)

	fmt.Println(bodyString)

}

func MultipleParameterValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	names := query["name"]

	fmt.Fprint(w, strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Albarra&name=Zikrillah", nil)

	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	bodyString := string(body)

	fmt.Println(bodyString)
}
