package errors

type dateFormatError struct {
}

func NewDateFormatError() *dateFormatError {
	return &dateFormatError{}
}

func (u *dateFormatError) Error() string {
	return "Date Format Error"
}
