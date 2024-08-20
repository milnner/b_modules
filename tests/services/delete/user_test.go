package delete

import (
	"database/sql"
	"testing"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/repositories"
	deleteService "github.com/milnner/b_modules/services/delete"
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
		// signUser := apptypes.SignUpUserType{
		// 	Name:       testCase[i].Name,
		// 	Surname:    testCase[i].Surname,
		// 	Professor:  1,
		// 	EntryDate:  testCase[i].EntryDate.Format(time.DateOnly),
		// 	BournDate:  testCase[i].BournDate.Format(time.DateOnly),
		// 	Email:      testCase[i].Email,
		// 	Password:   testCase[i].Hash,
		// 	Permission: testCase[i].Permision,
		// 	Sex:        testCase[i].Sex,
		// }
		var repositoryUserDelete *repositories.UserMySQLRepository
		if repositoryUserDelete, err = repositories.NewUserMySQLRepository(dbConn); err != nil {
			t.Fatal(err)
		}
		svc := deleteService.NewDeleteUserSvc(&testCase[i], repositoryUserDelete, config.Logger)
		svc.Run()
	}
}
