package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	firstname := r.PostForm.Get("first_name")
	lastname := r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstname, lastname)
}

func TestFormPost(t *testing.T) {
	payload := strings.NewReader("first_name=Albarra&last_name=Zikrillah")
	request := httptest.NewRequest(http.MethodPost, "/hi", payload)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	bytes, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello Albarra Zikrillah", string(bytes))
}
