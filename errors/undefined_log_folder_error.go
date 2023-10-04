package errors

type undefinedLogFolderError struct{}

func NewUndefinedLogFolderError() *undefinedLogFolderError {
	return &undefinedLogFolderError{}
}

func (u *undefinedLogFolderError) Error() string {
	return "Undefined folder log error"
}
