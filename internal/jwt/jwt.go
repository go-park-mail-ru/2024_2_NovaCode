package jwt

import (
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

func Generate(cfg *config.Config, user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(cfg.Auth.Jwt.Expire * time.Second).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(cfg.Auth.Jwt.Secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign jwt token: %w", err)
	}

	return tokenString, nil
}

func Verify(cfg *config.Config, accessToken string) error {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.Auth.Jwt.Secret), nil
	})
	if err != nil {
		return fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			expirationTime := time.Unix(int64(exp), 0)
			if expirationTime.Before(time.Now()) {
				return fmt.Errorf("token has expired")
			}
		} else {
			return fmt.Errorf("expiration claim not found")
		}

		return nil
	}

	return fmt.Errorf("invalid token")
}
