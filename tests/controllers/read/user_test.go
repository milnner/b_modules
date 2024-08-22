package read

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	readCtrl "github.com/milnner/b_modules/controllers/read"
	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/models"
	"github.com/milnner/b_modules/repositories"
	"github.com/milnner/b_modules/tests/config"
)

func TestReadUser(t *testing.T) {
	config.SetDBData()
	config.SetRootDatabaseConn()
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
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
		return
	}

	userRepo, err := repositories.NewUserMySQLRepository(dbConn)
	if err != nil {
		t.Fatal(err)
		return
	}
	ctrl := readCtrl.NewReadUserController(userRepo, config.Logger)

	var b []byte
	bd := bytes.NewBuffer(b)
	user := models.User{Email: config.UsersObjs[0].Email}
	err = json.NewEncoder(bd).Encode(user)
	wr := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "/", bd)

	ctrl.Handler(wr, req)

	err = json.NewDecoder(wr.Body).Decode(&user)
	if wr.Result().StatusCode != http.StatusOK {
		t.Errorf("Not Ok %v, status %v", user, wr.Result().StatusCode)
	}
}
