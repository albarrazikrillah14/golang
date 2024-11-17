package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	fmt.Fprint(w, contentType)
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "Medomeckz")
}

func TestHeader(t *testing.T) {
	t.Run("Request Header", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Add("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		RequestHeader(recorder, request)

		response := recorder.Result()
		bytes, _ := io.ReadAll(response.Body)
		assert.Equal(t, "application/json", string(bytes))
	})

	t.Run("Response Header", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		recorder := httptest.NewRecorder()

		ResponseHeader(recorder, request)

		poweredBy := recorder.Header().Get("X-Powered-By")

		assert.Equal(t, "Medomeckz", poweredBy)
	})
}
