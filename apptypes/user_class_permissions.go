package apptypes

type userClassPermissions struct {
	write Permission
	read  Permission
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
