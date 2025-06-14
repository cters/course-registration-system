package auth

import (
	"time"

	"github.com/QuanCters/backend/global"
	"github.com/gin-gonic/gin"
)

const (
	TokenCookieName = "auth_token"
)

func SetTokenCookie(c *gin.Context, token string) {
	timeEx := global.Config.JWT.JWT_EXPIRATION
	if timeEx == "" {
		timeEx = "24h"
	}

	expiration, err := time.ParseDuration(timeEx)
	if err != nil {
		expiration = 24 * time.Hour
	}

	c.SetCookie(
		TokenCookieName,
		token,
		int(expiration.Seconds()),
		"/",
		"",
		false,
		true,
	)
}


func GetTokenFromCookie(c *gin.Context) (string, error) {
	token, err := c.Cookie(TokenCookieName)
	if err != nil {
		return "", err
	}
	return token, nil
}


func ClearTokenCookie(c *gin.Context) {
	c.SetCookie(
		TokenCookieName,
		"",
		-1,
		"/",
		"",
		false,
		true,
	)
}