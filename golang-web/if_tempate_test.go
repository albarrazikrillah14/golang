package golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var tempates embed.FS

func IfTemplate(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(tempates, "templates/*.gohtml"))

	name := r.URL.Query().Get("name")
	user := &User{
		Title: "COI",
		Name:  name,
	}
	t.ExecuteTemplate(w, "if.gohtml", user)
}

func TestIfTemplate(t *testing.T) {
	t.Run("With name", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/?name=Albarra", nil)
		recorder := httptest.NewRecorder()

		IfTemplate(recorder, request)
		response, err := io.ReadAll(recorder.Result().Body)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(response))
	})
	t.Run("Without name", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		recorder := httptest.NewRecorder()

		IfTemplate(recorder, request)
		response, err := io.ReadAll(recorder.Result().Body)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(response))
	})
}
