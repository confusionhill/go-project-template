package authentication

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go-be-template/internal/model/dto/wrapper"
	"net/http"
	"strings"
)

type JwtClaims struct {
	jwt.Claims
}

func AuthorizationMiddleware(secret []byte) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")
			if tokenString == "" {
				return c.JSON(http.StatusUnauthorized, wrapper.BaseResponseDTO{
					Code:    http.StatusUnauthorized,
					Message: "missing bearer token",
				})
			}
			splitted := strings.Split(tokenString, " ")
			if len(splitted) != 2 || splitted[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, wrapper.BaseResponseDTO{
					Code:    http.StatusUnauthorized,
					Message: "bad authorization header",
				})
			}
			token, err := jwt.ParseWithClaims(splitted[1], &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
				return secret, nil
			})
			if err != nil {
				return c.JSON(http.StatusUnauthorized, wrapper.BaseResponseDTO{
					Code:    http.StatusUnauthorized,
					Message: "invalid token",
				})
			}

			if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
				c.Set("claims", claims)
				return next(c)
			}
			return c.JSON(http.StatusUnauthorized, wrapper.BaseResponseDTO{
				Code:    http.StatusInternalServerError,
				Message: "internal server error : cannot decode access token",
			})
		}
	}
}
