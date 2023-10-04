package errors

type unknownTypeOfSexError struct {
}

func NewUnknownTypeOfSexError() *unknownTypeOfSexError {
	return &unknownTypeOfSexError{}
}

func (u *unknownTypeOfSexError) Error() string {
	return "Unknown type of sex"
}
