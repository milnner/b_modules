package auth

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/milnner/b_modules/apptypes"
	authenticateCtrl "github.com/milnner/b_modules/controllers/auth"
	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/hasher"
	"github.com/milnner/b_modules/tests/config"
	"github.com/milnner/b_modules/tokens"
)

func TestAuthenticateUser(t *testing.T) {
	tkz := tokens.NewUserJWTokenizator(config.JwtSecretKey)
	var err error
	var dbConn *sql.DB

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
	user := apptypes.SignInUser{
		Email:    config.UsersObjs[0].Email,
		Password: config.UserMockPassword,
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("Post", "/user", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}

	ctrl := authenticateCtrl.NewAuthenticateUserController(config.DatabaseConn, config.Logger, tkz, hasher.NewBcryptHasher(), "mysql")

	handler := http.HandlerFunc(ctrl.Handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		var buff []byte
		t.Errorf("handler retornou o status errado: esperado %v mas obteve %v, body: %v", http.StatusOK, status, string(buff))
	}

	claims, err := tokens.NewUserJWTokenizator(config.JwtSecretKey).ValidateToken(rr.Body.String())

	if err != nil {
		t.Errorf("handler token invalido: token %v ;err: %v", rr.Body.String(), err)
	}

	if (((*claims)["email"]).(string)) != user.Email {
		t.Errorf("handler token invalido: email esperado %v ;obtido: %v", ((*claims)["Email"]).(string), user.Email)
	}
}
