package errors

type readRequestError struct {
}

func NewReadRequestError() *readRequestError {
	return &readRequestError{}
}

func (u *readRequestError) Error() string {
	return "Read request error"
}
