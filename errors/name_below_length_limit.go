package errors

type nameBelowLengthLimit struct {
}

func NewNameBelowLengthLimit() *nameBelowLengthLimit {
	return &nameBelowLengthLimit{}
}
func (u *nameBelowLengthLimit) Error() string {
	return "Name Below Length Limit [1<=|name|]"
}
