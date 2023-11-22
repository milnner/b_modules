package errors

type classEditorsNotFoundError struct {
}

func NewClassEditorsNotFoundError() *classEditorsNotFoundError {
	return &classEditorsNotFoundError{}
}

func (u *classEditorsNotFoundError) Error() string {
	return "Class Not Found Error"
}
