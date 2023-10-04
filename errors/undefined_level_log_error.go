package errors

type undefinedLevelLogError struct {
}

func NewUndefinedLevelLogError() *undefinedLevelLogError {
	return &undefinedLevelLogError{}
}

func (u *undefinedLevelLogError) Error() string {
	return "Undefined level log error"
}
