package auth

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateTokens(id uuid.UUID) (string, string, error) {
	log.Println(id.String())
	privateKey := []byte(os.Getenv("APISECRET"))

	refreshClaims := jwt.MapClaims{
		"authorized": true,
		"id":         id.String(),
		"exp":        time.Now().Add(time.Hour * 24 * time.Duration(365)).Unix(),
	}

	accessClaims := jwt.MapClaims{
		"authorized": true,
		"id":         id.String(),
		"exp":        time.Now().Add(time.Minute * time.Duration(10)).Unix(),
	}

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := refresh.SignedString(privateKey)
	if err != nil {
		log.Fatal(err)
		return "", "", err
	}

	access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := access.SignedString(privateKey)
	if err != nil {
		log.Fatal(err)
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func RefreshAccessToken(refreshToken string) (string, error) {
	privateKey := []byte(os.Getenv("APISECRET"))
	parsedToken, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Method.Alg())
		}
		return []byte(os.Getenv("APISECRET")), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return "", fmt.Errorf("invalid token claims")
	}
	exp, ok := claims["exp"].(float64)
	if !ok {
		return "", fmt.Errorf("missing or invalid 'exp' claim")
	}
	expTime := time.Unix(int64(exp), 0)

	if time.Now().After(expTime) {
		return "", fmt.Errorf("token is expired")
	}

	id := claims["id"]
	log.Println(id)
	accessClaims := jwt.MapClaims{
		"authorized": true,
		"id":         id,
		"exp":        time.Now().Add(time.Minute * time.Duration(10)).Unix(),
	}
	access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := access.SignedString(privateKey)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return accessToken, nil
}
