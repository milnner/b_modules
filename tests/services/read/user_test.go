package read

import (
	"database/sql"
	"testing"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/repositories"
	readService "github.com/milnner/b_modules/services/read"
	"github.com/milnner/b_modules/tests/config"
)

func TestCreateUser(t *testing.T) {
	config.SetRootDatabaseConn()
	var (
		err    error
		dbConn *sql.DB
	)

	testCase := config.UsersObjs

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM users WHERE 1"); err != nil {
			t.Fatal(err)
		}
	}()
	for i := 0; i < len(config.Users); i++ {
		if _, err = dbConn.Exec(config.Users[i]); err != nil {
			t.Fatal(err)
		}
	}
	for i := 0; i < len(testCase); i++ {
		var repositoryUserRead *repositories.UserMySQLRepository
		if repositoryUserRead, err = repositories.NewUserMySQLRepository(dbConn); err != nil {
			t.Fatal(err)
		}
		svc := readService.NewReadUserSvc(&testCase[i], repositoryUserRead)
		if svc.Run() != nil {
			t.Fatal(err)
		}
	}
}
