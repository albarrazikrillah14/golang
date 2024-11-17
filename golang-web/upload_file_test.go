package golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"testing"
)

//go:embed templates/*.gohtml
var tmpls embed.FS

func UploadForm(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(tmpls, "templates/*.gohtml"))
	t.ExecuteTemplate(w, "upload.form.gohtml", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// r.ParseMultipartForm(100 << 20)
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	fileDestination, err := os.Create("./resources" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	name := r.PostFormValue("name")

	fmt.Fprintln(w, name+fileHeader.Filename)
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
