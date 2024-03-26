package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"go-be-template/internal/model/config"
	"time"
)

type JwtClaims struct {
	jwt.Claims
}

func CreateAccessToken(cfg *config.JWTConfig, userId string) (string, error) {
	claims := JwtClaims{
		jwt.RegisteredClaims{
			Issuer:    "rest-be",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * cfg.Span)),
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Minute * cfg.Span)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        userId,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(cfg.AccessSecret))
}

func CreateRefreshToken(cfg *config.JWTConfig) {

}