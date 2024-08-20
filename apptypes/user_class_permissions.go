package apptypes

type UserClassPermission string

func (u UserClassPermission) Equals(o UserClassPermission) bool {
	return string(u) == string(o)
}

const (
	write UserClassPermission = "read"
	read  UserClassPermission = "write"
)

type userClassPermissions struct {
	write UserClassPermission
	read  UserClassPermission
}

var UserClassPermissions = userClassPermissions{
	write: write,
	read:  read,
}

func (u userClassPermissions) Read() string {
	return string(read)
}

func (u userClassPermissions) Write() string {
	return string(write)
}
