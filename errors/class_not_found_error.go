package errors

type classNotFoundError struct {
}

func NewClassNotFoundError() *classNotFoundError {
	return &classNotFoundError{}
}

func (u *classNotFoundError) Error() string {
	return "Class Not Found Error"
}
