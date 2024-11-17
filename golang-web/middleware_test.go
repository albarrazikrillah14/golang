package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (m *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before")
	m.Handler.ServeHTTP(w, r)
	fmt.Println("After")
}

type ErrorMiddleware struct {
	Handler http.Handler
}

func (e *ErrorMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error : %s", err)
		}
	}()

	e.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Executed")
		fmt.Fprint(w, "Hello Middleware")
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("Ups")
	})

	LogMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorMiddleware{
		Handler: LogMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
