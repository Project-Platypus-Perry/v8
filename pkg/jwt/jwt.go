package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/project-platypus-perry/v8/internal/config"
	"github.com/project-platypus-perry/v8/internal/constants"
)

type TokenType string

const (
	AccessToken  TokenType = "Access"
	RefreshToken TokenType = "Refresh"
)

type Claims struct {
	UserID         string             `json:"UserID"`
	Role           constants.UserRole `json:"Role"`
	Type           TokenType          `json:"Type"`
	OrganizationID string             `json:"OrganizationID"`
	jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken  string `json:"AccessToken"`
	RefreshToken string `json:"RefreshToken"`
}

func GenerateTokenPair(userID string, role constants.UserRole, organizationID string, cfg *config.JWTConfig) (*TokenPair, error) {
	// Generate Access Token
	accessToken, err := generateToken(userID, role, organizationID, AccessToken, cfg.AccessTokenSecret, cfg.AccessTokenExpiry)
	if err != nil {
		return nil, err
	}

	// Generate Refresh Token
	refreshToken, err := generateToken(userID, role, organizationID, RefreshToken, cfg.RefreshTokenSecret, cfg.RefreshTokenExpiry)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func generateToken(userID string, role constants.UserRole, organizationID string, tokenType TokenType, secret string, expiry time.Duration) (string, error) {
	claims := &Claims{
		UserID:         userID,
		Role:           role,
		OrganizationID: organizationID,
		Type:           tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string, secret string, expectedType TokenType) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if claims.Type != expectedType {
			return nil, errors.New("invalid token type")
		}
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
