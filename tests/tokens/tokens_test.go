package tokens

import (
	"testing"

	tokens "github.com/milnner/b_modules/tokens"
)

func TestGenerateToken(t *testing.T) {
	// Crie uma instância de JWTokenizator com uma chave secreta de teste
	tokenizator := tokens.NewJWTokenizator("teste-string")

	// Defina as reivindicações de teste
	claims := map[string]interface{}{
		"sub":   123,
		"email": "test@example.com",
	}

	// Gere um token
	tokenString, err := tokenizator.GenerateToken(claims)

	// Verifique se não há erros
	if err != nil {
		t.Errorf("Erro ao gerar token: %v", err)
	}

	// Verifique se o token gerado não está vazio
	if tokenString == "" {
		t.Errorf("Token gerado está vazio")
	}
}

func TestValidateToken(t *testing.T) {
	// Crie uma instância de JWTokenizator com uma chave secreta de teste
	tokenizator := tokens.NewJWTokenizator("teste-string")

	// Gere um token de teste
	claims := map[string]interface{}{
		"sub":   123,
		"email": "test@example.com",
	}

	tokenString, err := tokenizator.GenerateToken(claims)

	if err != nil {
		t.Errorf("Erro ao gerar token de teste: %v", err)
	}

	// Valide o token gerado
	user, err := tokenizator.ValidateToken(tokenString)

	// Verifique se não há erros
	if err != nil {
		t.Errorf("Erro ao validar token: %v", err)
	}

	// Verifique se as reivindicações do usuário são corretas
	if user == nil {
		t.Errorf("Reivindicações do usuário são nulas")
	} else {
		expectedUserID := 123
		expectedEmail := "test@example.com"

		if (*user)["Id"].(int) != expectedUserID {
			t.Errorf("ID do usuário incorreto. Esperado: %d, Obtido: %d", expectedUserID, (*user)["Id"].(int))
		}

		if (*user)["Email"].(string) != expectedEmail {
			t.Errorf("Email do usuário incorreto. Esperado: %s, Obtido: %s", expectedEmail, (*user)["Email"].(string))
		}
	}
}
