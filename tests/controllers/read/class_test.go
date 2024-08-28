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
	authSvc "github.com/milnner/b_modules/services/auth"
	"github.com/milnner/b_modules/tests/config"
	"github.com/milnner/b_modules/tokens"
)

func TestReadClass(t *testing.T) {
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
			Exec("DELETE FROM classes WHERE 1"); err != nil {
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

	for i := 0; i < len(config.Classes); i++ {
		if _, err = dbConn.
			Exec(config.Classes[i]); err != nil {
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

	classRepo, err := repositories.NewClassMySQLRepository(dbConn)
	if err != nil {
		t.Fatal(err)
		return
	}

	areaRepo, err := repositories.NewAreaMySQLRepository(dbConn)
	if err != nil {
		t.Fatal(err)
		return
	}

	ctrl := readCtrl.NewReadClassController(classRepo,
		areaRepo,
		tkz,
		config.Logger)

	var buff []byte
	body := bytes.NewBuffer(buff)
	class := models.Class{Id: config.ClassesObjs[0].Id, AreaId: config.ClassesObjs[0].AreaId}

	err = json.NewEncoder(body).Encode(class)

	wr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/", body)

	req.Header.Set("Authorization", "Bearer "+token)

	ctrl.Handler(wr, req)
	if wr.Result().StatusCode != http.StatusOK {
		t.Errorf("Not Ok, body  %v, status %v", wr.Body.String(), wr.Result().StatusCode)
	}
	class = models.Class{}
	if err := json.
		NewDecoder(wr.Result().Body).
		Decode(&class); err != nil {
		t.Fatal(err)
	}
	if !class.Equals(config.ClassesObjs[0]) {
		t.Errorf("Body  %v", wr.Body.String())
		t.Errorf("Classe need to be equal, but %v != %v, body %v ", class, config.ClassesObjs[0], wr.Body.String())
	}

	// teste para quando o token é invalido

	class = config.ClassesObjs[3]
	class.AreaId = config.AreasObjs[0].Id
	err = json.NewEncoder(body).Encode(class)

	wr = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/", body)
	req.Header.Set("Authorization", "Bearer invalid_token")
	ctrl.Handler(wr, req)
	if wr.Result().StatusCode != http.StatusUnauthorized {
		t.Errorf("Ok, but need to be Not Ok, body  %v", wr.Body.String())
	}

	// teste para quando o usuario não tem acesso a area
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

	class = models.Class{Id: config.ClassesObjs[1].Id, AreaId: class.AreaId}

	err = json.NewEncoder(body).Encode(class)

	wr = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/", body)
	req.Header.Set("Authorization", "Bearer "+token)
	ctrl.Handler(wr, req)
	if wr.Result().StatusCode != http.StatusUnauthorized {
		t.Errorf("Ok, but need to be Unaunthorized, status %v, body %v", wr.Result().StatusCode, wr.Body.String())
	}
}
