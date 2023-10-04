package errors

type notExistEnvironmentVariableError struct {
	varName string
}

func NewNotExistEnvironmentVariableError(varName string) *notExistEnvironmentVariableError {
	return &notExistEnvironmentVariableError{varName: varName}
}

func (u *notExistEnvironmentVariableError) Error() string {
	return "Not exist environment variable " + u.varName
}
