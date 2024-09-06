package create

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	createCtrl "github.com/milnner/b_modules/controllers/create"
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

		// if _, err = dbConn.
		// 	Exec("DELETE FROM class_takes_user WHERE 1"); err != nil {
		// 	t.Fatal(err)
		// }

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

	ctrl := createCtrl.NewCreateClassHasUserController(classRepo,
		areaRepo,
		tkz,
		config.Logger)

	var buff []byte
	body := bytes.NewBuffer(buff)
	userHasClassAccess := models.UserHasClassAccess{User: config.UsersObjs[0],
		Class: config.ClassesObjs[0]}

	err = json.NewEncoder(body).Encode(userHasClassAccess)

	wr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/", body)

	req.Header.Set("Authorization", "Bearer "+token)

	ctrl.Handler(wr, req)
	if wr.Result().StatusCode != http.StatusCreated {
		t.Errorf("Not Ok, body  %v, status %v", wr.Body.String(), wr.Result().StatusCode)
	}

	// teste para quando o token é invalido

	userHasClassAccess = models.UserHasClassAccess{User: config.UsersObjs[0],
		Class: config.ClassesObjs[0]}
	err = json.NewEncoder(body).Encode(userHasClassAccess)

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

	userHasClassAccess = models.UserHasClassAccess{User: user,
		Class: config.ClassesObjs[0]}

	err = json.NewEncoder(body).Encode(userHasClassAccess)

	wr = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/", body)
	req.Header.Set("Authorization", "Bearer "+token)
	ctrl.Handler(wr, req)
	if wr.Result().StatusCode != http.StatusUnauthorized {
		t.Errorf("Ok, but need to be Unaunthorized, status %v, body %v", wr.Result().StatusCode, wr.Body.String())
	}
}
