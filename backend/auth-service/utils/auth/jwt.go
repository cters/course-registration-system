package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JWTCfg struct {
	SecretKey string
	TokenExpiration time.Duration
}

type PayloadClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	UserID int32 `json:"user_id"`
}

type TokenManager struct {
	cfg *JWTCfg
}

func NewTokenManager(cfg *JWTCfg) *TokenManager {
	return &TokenManager{cfg: cfg}
}

func (tm *TokenManager) GenTokenJWT(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(tm.cfg.SecretKey))
}

func (tm *TokenManager) CreateToken(uuidToken, username string, userID int32) (string, error) {
	now := time.Now()
	expiresAt := now.Add(tm.cfg.TokenExpiration)

	return tm.GenTokenJWT(&PayloadClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "cters",
			Subject:   uuidToken,
		},
		Username: username,
		UserID:   userID,
	})
}

func (tm *TokenManager) ParseJwtTokenSubject(token string) (*PayloadClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &PayloadClaims{}, func(jwtToken *jwt.Token) (interface{}, error){
		return []byte(tm.cfg.SecretKey), nil
	}) 

	// Handle parsing error first
	if err != nil {
		return nil, fmt.Errorf("token parsing failed: %w", err)
	}

	// Validate token and claims
	if claims, ok := tokenClaims.Claims.(*PayloadClaims); ok && tokenClaims.Valid{
		return claims, nil
	}

	return nil, errors.New("invalid Token")
}