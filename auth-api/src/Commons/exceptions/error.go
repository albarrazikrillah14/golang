package exceptions

type ClientError struct {
	StatusCode int
	Message    string
	Name       string
}

func NewInvariantError(mesage string) *ClientError {
	return &ClientError{
		Name:       "InvariantError",
		StatusCode: 400,
		Message:    mesage,
	}
}

func NewNotFoundError(message string) *ClientError {
	return &ClientError{
		Name:       "NotFoundError",
		StatusCode: 404,
		Message:    message,
	}
}

func NewAuthenticationError(message string) *ClientError {
	return &ClientError{
		Name:       "AuthenticationError",
		StatusCode: 401,
		Message:    message,
	}
}
