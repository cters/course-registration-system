package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/QuanCters/backend/global"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type PayloadClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	UserID int32 `json:"user_id"`
}

func GenTokenJWT(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(global.Config.JWT.API_SECRET_KEY))
}

func CreateToken(uuidToken, username string, userID int32) (string, error) {
	timeEx := global.Config.JWT.JWT_EXPIRATION
	if timeEx == "" {
		timeEx = "1d"
	}

	expiration, err := time.ParseDuration(timeEx)
	if err != nil {
		return "", err
	}

	now := time.Now()
	expiresAt := now.Add(expiration)

	return GenTokenJWT(&PayloadClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID: uuid.New().String(),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt: jwt.NewNumericDate(now),
			Issuer: "cters",
			Subject: uuidToken,
		},
		Username: username,
		UserID: userID,
	})
}

func ParseJwtTokenSubject(token string) (*PayloadClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &PayloadClaims{}, func(jwtToken *jwt.Token) (interface{}, error){
		return []byte(global.Config.JWT.API_SECRET_KEY), nil
	}) 

	// Handle parsing error first
	if err != nil {
		return nil, fmt.Errorf("token parsing failed: %w", err)
	}

	// Validate token and claims
	if claims, ok := tokenClaims.Claims.(*PayloadClaims); ok && tokenClaims.Valid{
		return claims, nil
	}

	return nil, errors.New("Invalid Token")
}