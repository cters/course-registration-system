package auth

import (
	"time"

	"github.com/QuanCters/backend/global"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type PayloadClaims struct {
	jwt.RegisteredClaims
}

func GenTokenJWT(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(global.Config.JWT.API_SECRET_KEY))
}

func CreateToken(uuidToken string) (string, error) {
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
	})
}

func ParseJwtTokenSubject(token string) (*jwt.RegisteredClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &PayloadClaims{}, func(jwtToken *jwt.Token) (interface{}, error){
		return []byte(global.Config.JWT.API_SECRET_KEY), nil
	}) 

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*PayloadClaims); ok && tokenClaims.Valid{
			return &claims.RegisteredClaims, nil
		}
	}

	return nil, err
}

func VerifyTokenSubject(token string) (*jwt.RegisteredClaims, error) {
	claims, err := ParseJwtTokenSubject(token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}