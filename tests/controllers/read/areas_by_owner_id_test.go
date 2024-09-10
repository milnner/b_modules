package read

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

func TestReadAreasByOwnerId(t *testing.T) {
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

	ctrl := readCtrl.NewReadAreasByOwnerIdController(areaRepo, config.Logger, tkz)

	wr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/", nil)

	req.Header.Set("Authorization", "Bearer "+token)

	areas := []models.Area{}
	ctrl.Handler(wr, req)

	if wr.Result().StatusCode != http.StatusOK {
		t.Errorf("Not Ok, body  %v", wr.Body.String())
	}

	if err := json.NewDecoder(wr.Body).Decode(&areas); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(areas, config.AreasObjs) {
		t.Errorf("Not ok, areas: %v \nObjs:%v", areas, config.AreasObjs)
	}

	wr = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/", nil)
	req.Header.Set("Authorization", "Bearer invalid_token")
	ctrl.Handler(wr, req)
	if wr.Result().StatusCode != http.StatusUnauthorized {
		t.Errorf("Ok, but need to be Unauthorized, body  %v", wr.Body.String())
	}

	// Testando usuario sem areas
	// Preparando o jwt
	user = config.UsersObjs[2]

	if err = authSvc.
		NewAuthenticationSvc(&user,
			&token,
			tkz).
		Run(); err != nil {
		t.Fatal(err)
	}

	wr = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	ctrl.Handler(wr, req)
	fmt.Println(wr.Body.String())
	fmt.Println(wr.Result().StatusCode)
	if wr.Result().StatusCode != http.StatusNotFound {
		t.Errorf("Ok, but need to be No content, body  %v", wr.Body.String())
	}
}
