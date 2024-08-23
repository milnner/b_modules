package read

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	readCtrl "github.com/milnner/b_modules/controllers/read"
	"github.com/milnner/b_modules/repositories"
	authSvc "github.com/milnner/b_modules/services/auth"
	"github.com/milnner/b_modules/tokens"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/tests/config"
)

func TestReadUserIdsByAreaId(t *testing.T) {
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

	ctrl := readCtrl.NewReadUserIdsByAreaIdController(areaRepo, config.Logger, tkz)

	var buff []byte
	// testando para quem tem acesso a
	// area mas a area não tem usuarios relacionados

	body := bytes.NewBuffer(buff)
	area := config.AreasObjs[0]

	err = json.NewEncoder(body).Encode(area)

	wr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/", body)

	req.Header.Set("Authorization", "Bearer "+token)

	ctrl.Handler(wr, req)
	if wr.Result().StatusCode != http.StatusNotFound {
		t.Errorf("Not Ok, body  %v", wr.Body.String())
	}

	// testando para quem tem acesso a
	// area e a area tem usuarios relacionados
	var userIds []int

	users := config.UsersObjs[1:5]
	userObjsIds := []int{2, 3, 4, 5}

	for _, user := range users {
		user.Permision = "read"
		if err = areaRepo.InsertUser(&area, &user); err != nil {
			t.Fatal(err)
		}
	}

	area = config.AreasObjs[0]
	err = json.NewEncoder(body).Encode(area)
	wr = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/", body)
	req.Header.Set("Authorization", "Bearer "+token)
	ctrl.Handler(wr, req)

	if wr.Result().StatusCode != http.StatusOK {
		t.Errorf("Not Ok, body  %v", wr.Result().StatusCode)
	}

	userIds = []int{}
	json.NewDecoder(wr.Body).Decode(&userIds)

	if !reflect.DeepEqual(userIds, userObjsIds) {
		t.Errorf("Ok, but need to be Not Ok, body  %v", wr.Body.String())
	}

	// Testando token invalido
	// Preparando o jwt

	wr = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/", body)
	req.Header.Set("Authorization", "Bearer invalid_token")
	ctrl.Handler(wr, req)
	if wr.Result().StatusCode != http.StatusUnauthorized {
		t.Errorf("Ok, but need to be Not Ok, body  %v", wr.Body.String())
	}

	// Testando usuario fora da tabela de permissões
	// Preparando o jwt
	tkz = tokens.NewUserJWTokenizator(config.JwtSecretKey)
	user = config.UsersObjs[5]

	if err = authSvc.
		NewAuthenticationSvc(&user,
			&token,
			tkz).
		Run(); err != nil {
		t.Fatal(err)
	}

	area = config.AreasObjs[0]
	err = json.NewEncoder(body).Encode(area)

	wr = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/", body)
	req.Header.Set("Authorization", "Bearer "+token)
	ctrl.Handler(wr, req)

	if wr.Result().StatusCode != http.StatusNotFound {
		t.Errorf("Not Ok, body  %v", wr.Result().StatusCode)
	}

	// Testando usuario na tabela de permissões
	// Preparando o jwt
	tkz = tokens.NewUserJWTokenizator(config.JwtSecretKey)
	user = config.UsersObjs[4]

	if err = authSvc.
		NewAuthenticationSvc(&user,
			&token,
			tkz).
		Run(); err != nil {
		t.Fatal(err)
	}

	area = config.AreasObjs[0]
	err = json.NewEncoder(body).Encode(area)

	wr = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/", body)
	req.Header.Set("Authorization", "Bearer "+token)
	ctrl.Handler(wr, req)

	if wr.Result().StatusCode != http.StatusOK {
		t.Errorf("Not Ok, body  %v", wr.Result().StatusCode)
	}

	userIds = []int{}
	json.NewDecoder(wr.Body).Decode(&userIds)

	if !reflect.DeepEqual(userIds, userObjsIds) {
		t.Errorf("Ok, but need to be Not Ok, body  %v", wr.Body.String())
	}
}
