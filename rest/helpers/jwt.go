package helpers

import (
	"time"

	"example.com/rest/models"
	"github.com/golang-jwt/jwt/v5"
)

const SECRETKEY = "suppersecreat"

type PayloadToken struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`
}

type Claims struct {
	PayloadToken
	jwt.RegisteredClaims
}

func GenerateToken(data models.User) (string, error) {
	payload := Claims{
		PayloadToken: PayloadToken{
			UserId: data.ID.String(),
			Email:  data.Email,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(SECRETKEY))
}

func ValidateToken(tokenString string) (*PayloadToken, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRETKEY), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return &claims.PayloadToken, nil
	}

	return nil, jwt.ErrTokenInvalidClaims
}
