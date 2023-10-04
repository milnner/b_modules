package errors

type jWTInvalidTokenError struct {
}

func NewJWTInvalidTokenError() *jWTInvalidTokenError {
	return &jWTInvalidTokenError{}
}
func (u *jWTInvalidTokenError) Error() string {
	return "Invalid JWT"
}
