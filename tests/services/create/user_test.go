package create

import (
	"database/sql"
	"testing"
	"time"

	"github.com/milnner/b_modules/apptypes"
	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/hasher"
	"github.com/milnner/b_modules/repositories"
	createService "github.com/milnner/b_modules/services/create"
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

	for i := 0; i < len(testCase); i++ {
		signUser := apptypes.SignUpUser{
			Name:       testCase[i].Name,
			Surname:    testCase[i].Surname,
			Professor:  testCase[i].Professor,
			EntryDate:  testCase[i].EntryDate.Format(time.DateOnly),
			BournDate:  testCase[i].BournDate.Format(time.DateOnly),
			Email:      testCase[i].Email,
			Password:   testCase[i].Hash,
			Permission: testCase[i].Permision,
			Sex:        testCase[i].Sex,
		}
		var repositoryUserInsert *repositories.UserMySQLRepository
		if repositoryUserInsert, err = repositories.NewUserMySQLRepository(dbConn); err != nil {
			t.Fatal(err)
		}
		svc := createService.NewCreateUserSvc(signUser, repositoryUserInsert, hasher.NewBcryptHasher(), config.Logger)
		svc.Run()
	}
}
