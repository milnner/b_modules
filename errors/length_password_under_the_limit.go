package errors

type lengthPasswordUnderTheLimit struct {
}

func NewLengthPasswordUnderTheLimit() *lengthPasswordUnderTheLimit {
	return &lengthPasswordUnderTheLimit{}
}

func (u *lengthPasswordUnderTheLimit) Error() string {
	return "Length password under the limit [|password| >= 15]"
}
