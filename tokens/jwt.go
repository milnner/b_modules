package tokens

import (
	"github.com/golang-jwt/jwt/v4"
	modulesErrors "github.com/milnner/b_modules/errors"
)

type JWTokenizator struct {
	jwtSecretKey []byte
}

func NewJWTokenizator() *JWTokenizator {
	return &JWTokenizator{}
}

func (u JWTokenizator) GenerateToken(mapClaims map[string]interface{}) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(mapClaims))

	tokenString, err := claims.SignedString(u.jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u JWTokenizator) ValidateToken(tokenString string) (*map[string]interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return u.jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {

		user := make(map[string]interface{})
		user["Id"] = int((*claims)["sub"].(float64))
		user["Email"] = (*claims)["email"].(string)

		return &user, nil
	}

	return nil, modulesErrors.NewJWTInvalidTokenError()
}
