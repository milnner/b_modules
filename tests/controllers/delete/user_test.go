package delete

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	deleteCtrl "github.com/milnner/b_modules/controllers/delete"
	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/models"
	"github.com/milnner/b_modules/repositories"
	readSvc "github.com/milnner/b_modules/services/read"
	"github.com/milnner/b_modules/tests/config"
)

func TestDeleteUser(t *testing.T) {

	var (
		err    error
		dbConn *sql.DB
	)

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
		return
	}

	userRepo, err := repositories.NewUserMySQLRepository(dbConn)
	if err != nil {
		t.Fatal(err)
		return
	}

	ctrl := deleteCtrl.NewDeleteUserController(userRepo, config.Logger)

	var b []byte
	bd := bytes.NewBuffer(b)
	user := models.User{Id: config.UsersObjs[0].Id}
	err = json.NewEncoder(bd).Encode(user)
	wr := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "/", bd)

	ctrl.Handler(wr, req)

	if wr.Result().StatusCode != http.StatusOK {
		t.Errorf("Not Ok, status %v", wr.Result().StatusCode)
	}
	repo, err := repositories.NewUserMySQLRepository(dbConn)
	if err != nil {
		t.Errorf("Repo error %v", err.Error())
	}
	readService := readSvc.NewReadUserSvc(&user, repo)

	if err = readService.Run(); err != nil {
		t.Errorf("Service error %v", err.Error())
	}

	if user.Activated != 0 {
		t.Errorf("User is activated %v", user)
	}
}
