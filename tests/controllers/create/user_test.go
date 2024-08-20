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
	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/repositories"
	"github.com/milnner/b_modules/tests/config"
)

func TestInsertUser(t *testing.T) {
	config.SetRootDatabaseConn()
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

	user := apptypes.SignUpUser{
		Name:       "teste",
		Surname:    "Teste",
		Email:      "milnnernj@gmail.com",
		Professor:  1,
		EntryDate:  "",
		BournDate:  "2020-01-01",
		Permission: "write",
		Sex:        "male",
		Password:   "testetesteteste",
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Cria uma requisição HTTP fake
	req, err := http.NewRequest("Post", "/user", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
		return
	}
	userRepo, err := repositories.NewUserMySQLRepository(dbConn)
	if err != nil {
		t.Fatal(err)
		return
	}

	ctrl := createCtrl.NewCreateUserController(userRepo, config.Logger)
	// Cria um ResponseRecorder para capturar o response
	handler := http.HandlerFunc(ctrl.Handler)

	// Executa o handler com o request e o ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Verifica o status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retornou o status errado: esperado %v mas obteve %v", http.StatusOK, status)
	}

	// Verifica o conteúdo do body
	expected := `{}`
	if rr.Body.String() != expected {
		t.Errorf("handler retornou body inesperado: esperado %v mas obteve %v", expected, rr.Body.String())
	}
}
