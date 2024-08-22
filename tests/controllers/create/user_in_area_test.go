package create

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/milnner/b_modules/apptypes"
	createCtrl "github.com/milnner/b_modules/controllers/create"
	"github.com/milnner/b_modules/models"
	"github.com/milnner/b_modules/repositories"
	authSvc "github.com/milnner/b_modules/services/auth"
	"github.com/milnner/b_modules/tokens"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/tests/config"
)

func TestCreateUserInAreaAccess(t *testing.T) {
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
			Exec("DELETE FROM `user_has_area_access` WHERE 1"); err != nil {
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

	ctrl := createCtrl.NewCreateUserInAreaController(config.Logger, tkz, areaRepo)

	var buff []byte
	userInAreaAccess := models.UserHasAreaAccess{}
	body := bytes.NewBuffer(buff)
	userInAreaAccess.Area = config.AreasObjs[0]
	userInAreaAccess.User = config.UsersObjs[1]
	userInAreaAccess.User.Permision = apptypes.Permission(apptypes.UserAreaPermissions.Read())
	err = json.NewEncoder(body).Encode(userInAreaAccess)

	wr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/", body)

	req.Header.Set("Authorization", "Bearer "+token)

	ctrl.Handler(wr, req)
	if wr.Result().StatusCode != http.StatusCreated {
		t.Errorf("Not Ok, body  %v", wr.Body.String())
	}

	userInAreaAccess.User.Permision = ""
	if err = areaRepo.GetPermission(&userInAreaAccess.Area, &userInAreaAccess.User); err != nil ||
		userInAreaAccess.User.Permision != apptypes.
			Permission(apptypes.UserAreaPermissions.Read()) {
		t.Errorf("Permission is not insert, %v", err)
	}

	userInAreaAccess.Area = config.AreasObjs[1]
	userInAreaAccess.User = config.UsersObjs[1]
	userInAreaAccess.User.Permision = apptypes.Permission(apptypes.UserAreaPermissions.Read())
	err = json.NewEncoder(body).Encode(userInAreaAccess)

	wr = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/", body)
	req.Header.Set("Authorization", "Bearer invalid_token")
	ctrl.Handler(wr, req)
	if wr.Result().StatusCode == http.StatusOK {
		t.Errorf("Ok, but need to be Not Ok, body  %v", wr.Body.String())
	}

}
