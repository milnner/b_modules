package apptypes

type Permission string

const (
	write Permission = "write"
	read  Permission = "read"
)

func (u Permission) Equals(o Permission) bool {
	return string(u) == string(o)
}
