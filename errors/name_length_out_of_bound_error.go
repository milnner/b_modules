package errors

type nameLengthOutOfBoundError struct {
}

func NewNameLengthOutOfBoundError() *nameLengthOutOfBoundError {
	return &nameLengthOutOfBoundError{}
}

func (u nameLengthOutOfBoundError) Error() string {
	return "Name length out of bound error"
}
