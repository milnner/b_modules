package errors

type unreachableDatabaseStringsError struct {
	strings string
}

func NewUnreachableDatabaseStringsError(strings string) *unreachableDatabaseStringsError {
	return &unreachableDatabaseStringsError{strings: strings}
}

func (u *unreachableDatabaseStringsError) Error() string {
	return "unreachable Database Strings Error for" + u.strings
}
