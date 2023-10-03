package errors

type outOfFormatPasswordError struct {
}

func NewOutOfFormatPasswordError() *outOfFormatPasswordError {
	return &outOfFormatPasswordError{}
}

func (u *outOfFormatPasswordError) Error() string {
	return "Out ofF format password error"
}
