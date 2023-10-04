package errors

type emailFormatError struct {
}

func NewEmailFormatError() *emailFormatError {
	return &emailFormatError{}
}

func (u *emailFormatError) Error() string {
	return "Email format error"
}
