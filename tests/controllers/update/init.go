package update

import "github.com/milnner/b_modules/tests/config"

func init() {
	config.SetDBData()
	config.SetRootDatabaseConn()
}
