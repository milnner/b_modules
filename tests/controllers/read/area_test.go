package read

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	readCtrl "github.com/milnner/b_modules/controllers/read"
	"github.com/milnner/b_modules/models"
	"github.com/milnner/b_modules/repositories"
	authSvc "github.com/milnner/b_modules/services/auth"
	"github.com/milnner/b_modules/tokens"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/tests/config"
)

func TestReadArea(t *testing.T) {
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
			Exec("DELETE FROM user_has_area_access WHERE 1"); err != nil {
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
	for i := 0; i < len(config.Area); i++ {
		if _, err = dbConn.
			Exec(config.Area[i]); err != nil {
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

	ctrl := readCtrl.NewReadAreaController(areaRepo, config.Logger, tkz)

	var buff []byte
	body := bytes.NewBuffer(buff)
	area := models.Area{Id: config.AreasObjs[0].Id}
	area.OwnerId = 0

	err = json.NewEncoder(body).Encode(area)

	wr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/", body)

	req.Header.Set("Authorization", "Bearer "+token)

	ctrl.Handler(wr, req)
	if wr.Result().StatusCode != http.StatusOK {
		t.Errorf("Not Ok, body  %v", wr.Body.String())
	}
	if err := json.NewDecoder(wr.Body).Decode(&area); err != nil {
		t.Fatal(err)
	}
	if !area.Equals(config.AreasObjs[0]) {
		t.Errorf("Not ok, %v != %v", area, config.AreasObjs[0])
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

	// Testando usuario na tabela de permissões
	// Preparando o jwt
	tkz = tokens.NewUserJWTokenizator(config.JwtSecretKey)
	user = config.UsersObjs[2]

	if err = authSvc.
		NewAuthenticationSvc(&user,
			&token,
			tkz).
		Run(); err != nil {
		t.Fatal(err)
	}

	area = config.AreasObjs[1]
	area.OwnerId = 0
	err = json.NewEncoder(body).Encode(area)

	if err = areaRepo.InsertUser(&area, &user); err != nil {
		t.Fatal(err)
	}

	wr = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/", body)
	req.Header.Set("Authorization", "Bearer"+token)
	ctrl.Handler(wr, req)
	if wr.Result().StatusCode == http.StatusOK {
		t.Errorf("Ok, but need to be Not Ok, body  %v", wr.Body.String())
	}

}
