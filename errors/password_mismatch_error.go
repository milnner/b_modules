package errors

type passwordMismatchError struct {
}

func NewPasswordMismatchError() *passwordMismatchError {
	return &passwordMismatchError{}
}

func (u *passwordMismatchError) Error() string {
	return "Password mismatch error"
}
