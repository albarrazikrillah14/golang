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

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name != "" {
		fmt.Fprintf(w, "Hello %s", name)
	} else {
		fmt.Fprint(w, "Hello")
	}
}

func MultipleQueryParameter(w http.ResponseWriter, r *http.Request) {
	firstname := r.URL.Query().Get("first_name")
	lastname := r.URL.Query().Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstname, lastname)
}

func TestQueryParameter(t *testing.T) {

	t.Run("Without Query", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/hello", nil)
		recorder := httptest.NewRecorder()
		SayHello(recorder, request)

		response := recorder.Result()
		bytes, _ := io.ReadAll(response.Body)
		assert.Equal(t, "Hello", string(bytes))
	})

	t.Run("With Query", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/hello?name=Albarra", nil)
		recorder := httptest.NewRecorder()
		SayHello(recorder, request)
		response := recorder.Result()
		bytes, _ := io.ReadAll(response.Body)
		assert.Equal(t, "Hello Albarra", string(bytes))
	})

	t.Run("With Multiple Query Parameter", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/hello?first_name=Albarra&last_name=Zikrillah", nil)
		recorder := httptest.NewRecorder()

		MultipleQueryParameter(recorder, request)

		response := recorder.Result()
		bytes, _ := io.ReadAll(response.Body)
		assert.Equal(t, "Hello Albarra Zikrillah", string(bytes))
	})

	t.Run("With Multiple Parameter Values", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/hello?name=Albarra&name=Zikrillah", nil)
		recorder := httptest.NewRecorder()

		MultipleParameterValues(recorder, request)
		response := recorder.Result()
		bytes, _ := io.ReadAll(response.Body)
		assert.Equal(t, "Hello Albarra Zikrillah", string(bytes))
	})
}

func MultipleParameterValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]

	fmt.Fprint(w, "Hello ", strings.Join(names, " "))

}
