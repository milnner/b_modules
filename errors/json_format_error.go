package errors

type jSONFormatError struct {
}

func NewJSONFormatError() *jSONFormatError {
	return &jSONFormatError{}
}

func (u jSONFormatError) Error() string {
	return "JSON format error"
}
