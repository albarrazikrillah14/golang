package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "X-Medomeckz-Name",
		Value: r.URL.Query().Get("name"),
		Path:  "/",
	}

	http.SetCookie(w, cookie)

	fmt.Fprintln(w, "Success Create Cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-Medomeckz-Name")
	if err != nil {
		fmt.Fprintln(w, "No Cookie")
	} else {
		fmt.Fprintf(w, "Hello %s", cookie.Value)
	}
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
