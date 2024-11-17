package belajar_golang_web

import (
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)
	stripPrefix := http.StripPrefix("/static", fileServer)

	mux := http.NewServeMux()
	mux.Handle("/static/", stripPrefix)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
