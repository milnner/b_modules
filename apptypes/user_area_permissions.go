package apptypes

type userAreaPermissions struct {
	write Permission
	read  Permission
}

var UserAreaPermissions = userAreaPermissions{
	write: write,
	read:  read,
}

func (u userAreaPermissions) Read() string {
	return string(read)
}

func (u userAreaPermissions) Write() string {
	return string(write)
}
