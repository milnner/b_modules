package repositories

import (
	"database/sql"
	"testing"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/models"
	"github.com/milnner/b_modules/repositories"
	repositoryInterface "github.com/milnner/b_modules/repositories/interfaces"
	"github.com/milnner/b_modules/tests/config"
)

func TestUserMySQLRepositoryPolimorfism(t *testing.T) {
	var _ repositoryInterface.IUserRepository = &repositories.UserMySQLRepository{}
}

func TestGetUserById(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err := dbConn.Exec("DELETE FROM users WHERE 1"); err != nil {
			t.Fatal(err)
		}
	}()

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(config.Users); i++ {
		if _, err = dbConn.Exec(config.Users[i]); err != nil {
			t.Fatal(err)
		}
	}

	testCase := config.UsersObjs
	var repoUserSelect *repositories.UserMySQLRepository

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repoUserSelect, err = repositories.NewUserMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	var user models.User
	for _, tC := range testCase {
		user.Id = tC.Id
		if err = repoUserSelect.GetUserById(&user); err != nil ||
			user.Id != tC.Id ||
			user.Email != tC.Email ||
			user.Name != tC.Name ||
			user.Surname != tC.Surname ||
			user.Sex != tC.Sex ||
			user.Activated != tC.Activated {
			if err != nil {
				t.Errorf("Esperado ausencia de erro, mas %v", err)
				continue
			}
			t.Errorf("Esperado %v,\n mas %v\n", tC, user)
		}
	}
}

func TestGetUserByEmail(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err := dbConn.Exec("DELETE FROM users WHERE 1"); err != nil {
			t.Fatal(err)
		}
	}()

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(config.Users); i++ {
		if _, err = dbConn.Exec(config.Users[i]); err != nil {
			t.Fatal(err)
		}
	}

	testCase := config.UsersObjs
	var repoUserSelect *repositories.UserMySQLRepository

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repoUserSelect, err = repositories.NewUserMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	var user models.User
	for _, tC := range testCase {
		user.Email = tC.Email
		if err = repoUserSelect.GetUserByEmail(&user); err != nil ||
			user.Id != tC.Id ||
			user.Email != tC.Email ||
			user.Name != tC.Name ||
			user.Surname != tC.Surname ||
			user.Sex != tC.Sex ||
			user.Activated != tC.Activated {
			if err != nil {
				t.Errorf("Esperado ausencia de erro, mas %v", err)
				continue
			}
			t.Errorf("Esperado %v,\n mas %v\n", tC, user)
		}
	}
}

func TestGetUsersByIds(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err := dbConn.Exec("DELETE FROM users WHERE 1"); err != nil {
			t.Fatal(err)
		}
	}()
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.Users); i++ {
		if _, err = dbConn.Exec(config.Users[i]); err != nil {
			t.Fatal(err)
		}
	}

	var testCases models.Users = config.UsersObjs

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}
	var repoUserSelect *repositories.UserMySQLRepository
	if repoUserSelect, err = repositories.NewUserMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	var users []models.User
	for _, v := range testCases {
		users = append(users, models.User{Id: v.Id})
	}
	if err = repoUserSelect.GetUsersByIds(users); err != nil {
		t.Error(err)
	}

	if len(testCases) != len(users) {
		t.Errorf("Esperado capturar %v areas, mas recebeu %v areas\n", len(testCases), len(users))
	} else {
		for i, tC := range testCases {
			if tC.Id != users[i].Id ||
				tC.Name != users[i].Name ||
				tC.Surname != users[i].Surname ||
				tC.Email != users[i].Email ||
				tC.Hash != users[i].Hash ||
				tC.Activated != users[i].Activated {
				t.Errorf("Esperado %v, mas %v", tC, users[i])
			}
		}
	}
}

func TestUpdateUser(t *testing.T) {
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
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.Users); i++ {
		if _, err = dbConn.Exec(config.Users[i]); err != nil {
			t.Fatal(err)
		}
	}
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	var repositoryUserInsert *repositories.UserMySQLRepository
	if repositoryUserInsert, err = repositories.NewUserMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	for _, tC := range testCase {
		tC.Hash = "milnner"
		if err = repositoryUserInsert.Update(&tC); err != nil {
			t.Error(err)
		}
	}
}

func TestInsertUser(t *testing.T) {
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
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	var repositoryUserInsert *repositories.UserMySQLRepository
	if repositoryUserInsert, err = repositories.NewUserMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	for _, tC := range testCase {
		if err = repositoryUserInsert.Insert(&tC); err != nil {
			t.Error(err)
		}
	}
}

func TestDeleteUser(t *testing.T) {
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
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.Users); i++ {
		if _, err = dbConn.Exec(config.Users[i]); err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetDelete(), "mysql"); err != nil {
		t.Fatal(err)
	}

	var repositoryUserInsert *repositories.UserMySQLRepository
	if repositoryUserInsert, err = repositories.NewUserMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	for _, tC := range testCase {
		if err = repositoryUserInsert.Delete(&tC); err != nil {
			t.Error(err)
		}
	}
}
