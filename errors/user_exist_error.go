package errors

type userExistError struct {
}

func NewUserExistError() *userExistError {
	return &userExistError{}
}

func (u *userExistError) Error() string {
	return "User exist error"
}
