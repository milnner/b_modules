package create

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	createCtrl "github.com/milnner/b_modules/controllers/create"
	"github.com/milnner/b_modules/repositories"
	authSvc "github.com/milnner/b_modules/services/auth"
	"github.com/milnner/b_modules/tokens"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/tests/config"
)

func TestCreateArea(t *testing.T) {
	config.SetRootDatabaseConn()
	var err error
	var dbConn *sql.DB

	defer func() {
		if err = database.
			InitDatabaseConn(&dbConn,
				config.
					DatabaseConn.
					User.GetDelete(),
				"mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.
			Exec("DELETE FROM area WHERE 1"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.
			Exec("DELETE FROM users WHERE 1"); err != nil {
			t.Fatal(err)
		}

	}()

	if err = database.
		InitDatabaseConn(&dbConn,
			config.DatabaseConn.
				User.GetInsert(),
			"mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.Users); i++ {
		if _, err = dbConn.
			Exec(config.Users[i]); err != nil {
			t.Fatal(err)
		}
	}

	// Preparando o jwt
	var token string
	tkz := tokens.NewUserJWTokenizator(config.JwtSecretKey)
	user := config.UsersObjs[0]

	if err = authSvc.
		NewAuthenticationSvc(&user,
			&token,
			tkz).
		Run(); err != nil {
		t.Fatal(err)
	}

	if err = database.
		InitDatabaseConn(&dbConn,
			config.DatabaseConn.
				User.GetInsert(),
			"mysql"); err != nil {
		t.Fatal(err)
		return
	}
	areaRepo, err := repositories.NewAreaMySQLRepository(dbConn)
	if err != nil {
		t.Fatal(err)
		return
	}

	ctrl := createCtrl.NewCreateAreaController(areaRepo, config.Logger, tkz)

	var buff []byte
	body := bytes.NewBuffer(buff)
	area := config.AreasObjs[0]
	area.OwnerId = 0

	err = json.NewEncoder(body).Encode(area)

	wr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/", body)

	req.Header.Set("Authorization", "Bearer "+token)

	ctrl.Handler(wr, req)
	if wr.Result().StatusCode != http.StatusCreated {
		t.Errorf("Not Ok, body  %v", wr.Body.String())
	}

	area = config.AreasObjs[1]
	area.OwnerId = 0
	err = json.NewEncoder(body).Encode(area)

	wr = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/", body)
	req.Header.Set("Authorization", "Bearer invalid_token")
	ctrl.Handler(wr, req)
	if wr.Result().StatusCode == http.StatusOK {
		t.Errorf("Ok, but need to be Not Ok, body  %v", wr.Body.String())
	}

}
