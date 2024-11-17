package exceptions

import "net/http"

type Error struct {
	StatusCode int
	Name       string
	Message    string
}

func InvariantError(message string) Error {
	return Error{
		StatusCode: http.StatusBadRequest,
		Name:       "InvariantErrror",
		Message:    message,
	}
}

func NotFoundError(message string) Error {
	return Error{
		StatusCode: http.StatusNotFound,
		Name:       "NotFoundError",
		Message:    message,
	}
}
