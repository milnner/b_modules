package tokens

import (
	"net/http"
	"strings"
)

type IJWTokenizator interface {
	GenerateToken(mapClaims map[string]interface{}) (string, error)
	ValidadeToken(tokenString string) (*map[string]interface{}, error)
}

func ExtractTokenFromRequest(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}
