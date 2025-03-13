package jwt

import (
	"backend/internal/infra/env"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type JWTI interface {
	GenerateToken(userID uuid.UUID, isAdmin bool) (string, error)
	ValidateToken(tokenString string) (uuid.UUID, bool, error)
}

type JWT struct {
	secretKey   string
	expiredTime time.Duration
}

type Claims struct {
	ID      uuid.UUID
	IsAdmin bool
	jwt.RegisteredClaims
}

func NewJWT(env *env.Env) JWTI {
	secretKey := env.JWTSecret
	expiredTime := env.JWTExpired

	return &JWT{
		secretKey:   secretKey,
		expiredTime: expiredTime,
	}
}

func (j *JWT) GenerateToken(userID uuid.UUID, isAdmin bool) (string, error) {

	claims := Claims{
		ID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.expiredTime * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWT) ValidateToken(tokenString string) (uuid.UUID, bool, error) {
	var claims Claims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return uuid.Nil, false, err
	}

	if !token.Valid {
		return uuid.Nil, false, errors.New("invalid token")
	}

	userID := claims.ID
	isAdmin := claims.IsAdmin

	return userID, isAdmin, nil

}
