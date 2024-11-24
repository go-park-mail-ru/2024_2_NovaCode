package utils

import (
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(cfg *config.JwtConfig, user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.UserID.String(),
		"role":    user.Role,
		"exp":     time.Now().Add(cfg.Expire * time.Second).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.Secret))
}

// func VerifyJWT(cfg *config.JwtConfig, accessToken string) (uuid.UUID, error) {
// 	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(cfg.Secret), nil
// 	})
// 	if err != nil {
// 		return uuid.Nil, fmt.Errorf("failed to parse token: %w", err)
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || !token.Valid {
// 		return uuid.Nil, fmt.Errorf("invalid token")
// 	}

// 	if exp, ok := claims["exp"].(float64); ok {
// 		if time.Unix(int64(exp), 0).Before(time.Now()) {
// 			return uuid.Nil, fmt.Errorf("token has expired")
// 		}
// 	} else {
// 		return uuid.Nil, fmt.Errorf("expiration claim not found")
// 	}

// 	userIDStr, ok := claims["user_id"].(string)
// 	if !ok {
// 		return uuid.Nil, fmt.Errorf("user_id claim not found")
// 	}

// 	userID, err := uuid.Parse(userIDStr)
// 	if err != nil {
// 		return uuid.Nil, fmt.Errorf("invalid user ID format in token: %w", err)
// 	}

// 	return userID, nil
// }

func VerifyJWT(cfg *config.JwtConfig, accessToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.Secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Verify expiration
	if exp, ok := claims["exp"].(float64); ok {
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			return nil, fmt.Errorf("token has expired")
		}
	} else {
		return nil, fmt.Errorf("expiration claim not found")
	}

	return claims, nil
}
