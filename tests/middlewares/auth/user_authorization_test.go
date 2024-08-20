package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	authMiddleware "github.com/milnner/b_modules/middlewares/auth"

	authSvc "github.com/milnner/b_modules/services/auth"
	"github.com/milnner/b_modules/tests/config"
	"github.com/milnner/b_modules/tokens"
)

func TestUserAuthenticationMiddleware(t *testing.T) {
	config.SetDBData()
	config.SetRootDatabaseConn()
	var err error
	var token string

	// Supondo que o primeiro usuário seja utilizado para o teste
	user := config.UsersObjs[0]
	tkz := tokens.NewUserJWTokenizator(config.JwtSecretKey)

	middleware := authMiddleware.NewUserAuthorizationMiddleware(tkz)

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	handler := middleware.Handler(testHandler)

	server := httptest.NewServer(handler)
	defer server.Close()

	authenticationSvc := authSvc.NewAuthenticationSvc(&user, &token, tkz)
	if err = authenticationSvc.Run(); err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", server.URL, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status AlreadyReported, got %v", resp.StatusCode)
	}

	req, err = http.NewRequest("GET", server.URL, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer invalid_token")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	// Verificando se o status retornado é 401 Unauthorized
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status Unauthorized, got %v", resp.StatusCode)
	}
}
