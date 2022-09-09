package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func ValidateToken(ctx *gin.Context) {
	tokenString, err := validateAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, "Invalid token signature")
			return
		}
		ctx.JSON(http.StatusBadRequest, "Invalid token")
		return
	}

	if !token.Valid {
		ctx.JSON(http.StatusUnauthorized, "Invalid token")
		return
	}

	ctx.Next()
}

func validateAuthorization(ctx *gin.Context) (string, error) {
	bearerToken := ctx.GetHeader("Authorization")
	if bearerToken == "" {
		return "", errors.New("authorization header is required")
	}

	token := strings.Fields(bearerToken)
	return token[1], nil
}
