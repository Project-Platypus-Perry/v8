package utils

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GeneratePassword generates a secure random password of specified length
func GeneratePassword(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

// GeneratePasswordResetToken generates a JWT token for password reset
func GeneratePasswordResetToken(userID string, secret string, expiryHours int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"purpose": "password_reset",
		"exp":     time.Now().Add(time.Hour * time.Duration(expiryHours)).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ValidatePasswordResetToken validates the password reset token
func ValidatePasswordResetToken(tokenString string, secret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if purpose, ok := claims["purpose"].(string); ok && purpose == "password_reset" {
			if userID, ok := claims["user_id"].(string); ok {
				return userID, nil
			}
		}
	}

	return "", jwt.ErrInvalidKey
}
