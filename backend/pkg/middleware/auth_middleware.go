package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/golang-jwt/jwt/v4" 
)

func JWTMiddleware(requiredRole string) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("your_secret_key"),
		Claims:     jwt.MapClaims{},
		SuccessHandler: func(c echo.Context) {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			userRole := claims["role"].(string)

			if userRole != requiredRole {
				c.Error(echo.NewHTTPError(http.StatusForbidden, "Access denied"))
				return
			}
		},
	})
}