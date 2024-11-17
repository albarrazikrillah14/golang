package test

import (
	"medomeckz/auth-api/src/Commons/exceptions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientError(t *testing.T) {
	clientError := exceptions.ClientError{
		Name:       "ClientError",
		StatusCode: 400,
		Message:    "error",
	}

	assert.NotNil(t, clientError)
	assert.Equal(t, "ClientError", clientError.Name)
	assert.Equal(t, 400, clientError.StatusCode)
	assert.Equal(t, "error", clientError.Message)
}

func TestInvariantError(t *testing.T) {
	err := exceptions.NewInvariantError("terjadi kesalahan")
	assert.NotNil(t, err)
	assert.Equal(t, "InvariantError", err.Name)
	assert.Equal(t, 400, err.StatusCode)
	assert.Equal(t, "terjadi kesalahan", err.Message)
}

func TestNotFoundError(t *testing.T) {
	err := exceptions.NewNotFoundError("data tidak ditemukan")
	assert.NotNil(t, err)
	assert.Equal(t, "NotFoundError", err.Name)
	assert.Equal(t, 404, err.StatusCode)
	assert.Equal(t, "data tidak ditemukan", err.Message)
}

func TestAuthenticationError(t *testing.T) {
	err := exceptions.NewAuthenticationError("token tidak valid")
	assert.NotNil(t, err)
	assert.Equal(t, "AuthenticationError", err.Name)
	assert.Equal(t, 401, err.StatusCode)
	assert.Equal(t, "token tidak valid", err.Message)
}
