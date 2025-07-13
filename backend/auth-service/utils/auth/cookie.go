package auth

import (
	"time"

	"github.com/gin-gonic/gin"
)

const (
	TokenCookieName = "auth_token"
)

type CookieManager struct{
	JWTExpiration time.Duration
}


func NewCookieManager(jwtExpiration time.Duration) *CookieManager {
	return &CookieManager{
		JWTExpiration: jwtExpiration,
	}
}


func (cm *CookieManager) SetTokenCookie(c *gin.Context, token string) {
	c.SetCookie(
		TokenCookieName,
		token,
		int(cm.JWTExpiration.Seconds()),
		"/",
		"",
		false,
		true,
	)
}


func (cm *CookieManager) GetTokenFromCookie(c *gin.Context) (string, error) {
	return c.Cookie(TokenCookieName) 
}


func (cm *CookieManager) ClearTokenCookie(c *gin.Context) {
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