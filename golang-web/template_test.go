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

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	// t, err := template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	t := template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(w, "SIMPLE", "Hello HTML Template")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	SimpleHTML(recorder, request)
	response, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(response))
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/simple.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "simple.gohtml", "Simple")
}

func TestSimpleHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	SimpleHTMLFile(recorder, request)
	response, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(response))
}

func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

func TestSimpleHTMLDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	TemplateDirectory(recorder, request)
	response, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(response))
}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello Sir")
}

func TestSimpleHTMLDEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	TemplateEmbed(recorder, request)
	response, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(response))
}

type User struct {
	Title string
	Name  string
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	user := &User{
		Title: "Halo",
		Name:  "Albarra Zikrillah",
	}
	t.ExecuteTemplate(w, "name.gohtml", user)
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	TemplateDataStruct(recorder, request)

	response, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(response))
}
