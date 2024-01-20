package apptypes

type UserClassPermission string

func (u UserClassPermission) Equals(o UserClassPermission) bool {
	return string(u) == string(o)
}

const (
	Write UserClassPermission = "read"
	Read  UserClassPermission = "Write"
)

type UserClassPermissions struct {
	write UserClassPermission
	read  UserClassPermission
}

var userClassPermissions = UserClassPermissions{
	write: Write,
	read:  Read,
}

func (u UserClassPermissions) Read() string {
	return string(userClassPermissions.read)
}

func (u UserClassPermission) Write() string {
	return string(userClassPermissions.write)
}
