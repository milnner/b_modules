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
	"github.com/milnner/b_modules/models"
	"github.com/milnner/b_modules/repositories"
	authSvc "github.com/milnner/b_modules/services/auth"
	"github.com/milnner/b_modules/tokens"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/tests/config"
)

func TestReadClassIdsByAreaId(t *testing.T) {
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

	if err = database.
		InitDatabaseConn(&dbConn,
			config.DatabaseConn.
				User.GetSelect(),
			"mysql"); err != nil {
		t.Fatal(err)
		return
	}

	areaRepo, err := repositories.NewAreaMySQLRepository(dbConn)
	if err != nil {
		t.Fatal(err)
		return
	}
	classRepo, err := repositories.NewClassMySQLRepository(dbConn)
	if err != nil {
		t.Fatal(err)
		return
	}

	ctrl := readCtrl.NewReadClassIdsByAreaIdController(classRepo, areaRepo, config.Logger, tkz)
	var buff []byte
	body := bytes.NewBuffer(buff)
	area := models.Area{Id: config.AreasObjs[0].Id}
	area.OwnerId = 0

	err = json.NewEncoder(body).Encode(area)

	wr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/", body)

	req.Header.Set("Authorization", "Bearer "+token)

	classIds := []int{}
	ctrl.Handler(wr, req)

	if wr.Result().StatusCode != http.StatusOK {
		t.Errorf("Not Ok, body  %v, status %v", wr.Body.String(), wr.Result().StatusCode)
	}

	if err := json.NewDecoder(wr.Body).Decode(&classIds); err != nil {
		t.Fatal(err)
	}

	areasObjsIds := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(classIds, areasObjsIds) {
		t.Errorf("Not ok, areas: %v \nObjs:%v", classIds, areasObjsIds)
	}
	classIds = []int{}

	wr = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/", body)
	req.Header.Set("Authorization", "Bearer invalid_token")
	ctrl.Handler(wr, req)
	if wr.Result().StatusCode != http.StatusUnauthorized {
		t.Errorf("Ok, but need to be Unauthorized, body  %v", wr.Body.String())
	}

	// Testando usuario sem autorização
	// Preparando o jwt
	user = config.UsersObjs[2]

	if err = authSvc.
		NewAuthenticationSvc(&user,
			&token,
			tkz).
		Run(); err != nil {
		t.Fatal(err)
	}

	area = models.Area{Id: config.AreasObjs[0].Id}
	area.OwnerId = 0

	err = json.NewEncoder(body).Encode(area)

	wr = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/", body)

	req.Header.Set("Authorization", "Bearer "+token)
	ctrl.Handler(wr, req)
	if wr.Result().StatusCode != http.StatusUnauthorized {
		t.Errorf("Ok, but need to be No content, body  %v, status %v", wr.Body.String(), wr.Result().StatusCode)
	}
	if len(classIds) > 0 {
		t.Errorf("Not ok, areasIds: %v ", classIds)
	}
}
