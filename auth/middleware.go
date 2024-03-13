package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := tokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func tokenValid(c *gin.Context) error {
	token := ExtractToken(c)

	parsedToken, err := parseToken(token)
	if err != nil {
		return err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return fmt.Errorf("invalid token claims")
	}
	exp, ok := claims["exp"].(float64)
	if !ok {
		return fmt.Errorf("missing or invalid 'exp' claim")
	}
	expTime := time.Unix(int64(exp), 0)

	if time.Now().After(expTime) {
		return fmt.Errorf("token is expired")
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func parseToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Method.Alg())
		}
		return []byte(os.Getenv("APISECRET")), nil
	})

	if err != nil {
		return parsedToken, err
	}

	return parsedToken, nil
}

func ExtractIdFromContext(c *gin.Context) (uuid.UUID, error) {
	token := ExtractToken(c)
	return ExtractIdFromToken(token)
}

func ExtractIdFromToken(token string) (uuid.UUID, error) {
	var id uuid.UUID
	parsedToken, err := parseToken(token)
	if err != nil {
		return id, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return id, fmt.Errorf("invalid token claims")
	}
	idStr, ok := claims["id"].(string)
	if !ok {
		return id, fmt.Errorf("id is null or not string")
	}
	id, err = uuid.Parse(idStr)
	if err != nil {
		return id, err
	}
	return id, nil
}
