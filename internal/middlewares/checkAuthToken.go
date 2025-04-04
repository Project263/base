package middlewares

import (
	"base/config"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type AuthMiddleware struct {
	cfg *config.Config
}

func NewAuthMiddleware(cfg *config.Config) *AuthMiddleware {
	return &AuthMiddleware{cfg: cfg}
}

func (m *AuthMiddleware) CheckAuthToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Cookies()
		var authToken string

		for _, cookie := range token {
			if cookie.Name == "token" {
				authToken = cookie.Value
				break
			}
		}

		if authToken == "" {
			logrus.Error("Auth token not found")
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized: Token not found"})
		}

		if !isValidToken(m.cfg, authToken) {
			logrus.Error("Invalid token")
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized: Invalid token"})
		}

		return next(c)
	}
}

func isValidToken(cfg *config.Config, tokenString string) bool {
	secret := cfg.SECRET
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false
	}

	if claims == nil {
		return false
	}

	return true
}
