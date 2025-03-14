package jwt

import (
	"backend/internal/infra/env"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"strconv"
	"time"
)

type JWTI interface {
	GenerateToken(userID uuid.UUID, isAdmin bool, role string) (string, error)
	ValidateToken(tokenString string) (uuid.UUID, bool, string, error)
}

type JWT struct {
	secretKey   string
	expiredTime time.Duration
}

type Claims struct {
	ID      uuid.UUID
	IsAdmin bool
	Role    string
	jwt.RegisteredClaims
}

func NewJWT(env *env.Env) JWTI {
	secretKey := env.JWTSecret
	expiredTime, err := strconv.Atoi(os.Getenv("JWT_EXPIRED"))
	if err != nil {
		return nil
	}

	return &JWT{
		secretKey:   secretKey,
		expiredTime: time.Duration(expiredTime) * time.Hour,
	}
}

func (j *JWT) GenerateToken(userID uuid.UUID, isAdmin bool, role string) (string, error) {

	claims := Claims{
		ID:      userID,
		IsAdmin: isAdmin,
		Role:    role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.expiredTime)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWT) ValidateToken(tokenString string) (uuid.UUID, bool, string, error) {
	var claims Claims

	token, err := jwt.ParseWithClaims(tokenString, &claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.secretKey), nil
		})

	if err != nil {
		return uuid.Nil, false, "", err
	}

	if !token.Valid {
		return uuid.Nil, false, "", errors.New("invalid token")
	}

	userID := claims.ID
	isAdmin := claims.IsAdmin
	role := claims.Role

	return userID, isAdmin, role, nil

}
