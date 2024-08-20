package tokens

import (
	"github.com/golang-jwt/jwt/v4"
	modulesErrors "github.com/milnner/b_modules/errors"
)

type UserJWTokenizator struct {
	jwtSecretKey []byte
}

func NewUserJWTokenizator(jwtSecretKey string) *UserJWTokenizator {
	return &UserJWTokenizator{jwtSecretKey: []byte(jwtSecretKey)}
}

// ValidadeToken implements IJWTokenizator.
func (u *UserJWTokenizator) ValidateToken(tokenString string) (*map[string]interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return u.jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {

		user := make(map[string]interface{})
		user["sub"] = int((*claims)["sub"].(float64))
		user["email"] = (*claims)["email"].(string)

		return &user, nil
	}

	return nil, modulesErrors.NewJWTInvalidTokenError()
}

// GenerateToken(mapClaims map[string]interface{}) (string, error)
func (u UserJWTokenizator) GenerateToken(mapClaims map[string]interface{}) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(mapClaims))

	tokenString, err := claims.SignedString(u.jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
