package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-Medomeckz-Name"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"
	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Success Create Cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := r.Cookies()
	fmt.Fprint(w, cookie)
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
