package util

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecret string = os.Getenv("JWT_SECRET")

type userClaims struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type JWTClaims struct {
	userClaims
	jwt.RegisteredClaims
}

func CreateJWT(fullName, email, role string) (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	sub := strings.ToLower(strings.Split(fullName, " ")[0])

	claims := JWTClaims{
		userClaims{
			FullName: fullName,
			Email:    email,
			Role:     role,
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "CloServer",
			Subject:   sub,
			ID:        id.String(),
			Audience: jwt.ClaimStrings{
				"api",
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	claims := new(JWTClaims)

	return jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
}
