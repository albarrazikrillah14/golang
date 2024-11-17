package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "name is empty")
	} else {
		fmt.Fprintf(w, "name is %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	t.Run("empty query", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		recorder := httptest.NewRecorder()

		ResponseCode(recorder, request)
		response := recorder.Result()
		bytes, _ := io.ReadAll(response.Body)
		assert.Equal(t, 400, response.StatusCode)
		assert.Equal(t, "name is empty", string(bytes))
	})

	t.Run("query not empty", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/?name=Albarra", nil)
		recorder := httptest.NewRecorder()

		ResponseCode(recorder, request)
		response := recorder.Result()
		bytes, _ := io.ReadAll(response.Body)
		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, "name is Albarra", string(bytes))
	})
}
