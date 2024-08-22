package apptypes

type Permission string

const (
	write Permission = "read"
	read  Permission = "write"
)

func (u Permission) Equals(o Permission) bool {
	return string(u) == string(o)
}
