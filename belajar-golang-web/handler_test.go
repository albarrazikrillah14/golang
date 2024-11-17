package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "Hello World")
	})

	mux.HandleFunc("/hi", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "hi")
	})

	mux.HandleFunc("/images/", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "images")
	})

	mux.HandleFunc("/images/thumbnails", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "thumbnails")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
