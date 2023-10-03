package errors

type surnameLengthOutOfBoundError struct {
}

func NewSurnameLengthOutOfBoundError() *surnameLengthOutOfBoundError {
	return &surnameLengthOutOfBoundError{}
}

func (u surnameLengthOutOfBoundError) Error() string {
	return "Surname length out of bound error"
}
