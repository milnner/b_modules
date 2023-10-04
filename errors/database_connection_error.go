package errors

type databaseConnectionError struct {
}

func NewDatabaseConnectionError() *databaseConnectionError {
	return &databaseConnectionError{}
}

func (u *databaseConnectionError) Error() string {
	return "Database connection error"
}
