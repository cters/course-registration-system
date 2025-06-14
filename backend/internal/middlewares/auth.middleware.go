package middlewares

import (
	"context"
	"net/http"

	"github.com/QuanCters/backend/internal/utils/auth"
	"github.com/QuanCters/backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type contextKey string
const (
	SubjectUUIDKey contextKey = "subjectUUID"
	Username contextKey = "username"
	UserID contextKey = "user_id"
)


func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken, valid := auth.GetTokenFromCookie(c)
		if valid != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": response.ErrCodeAuthFailed,
				"message": "Authentication required",
			})
			c.Abort()
			return
		}

		claims, err := auth.ParseJwtTokenSubject(jwtToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    response.ErrCodeAuthFailed,
				"message": "Invalid token: " + err.Error(),
			})
			c.Abort()
			return
		}

		ctx := context.WithValue(c.Request.Context(), SubjectUUIDKey, claims.RegisteredClaims.Subject)
		ctx = context.WithValue(ctx, Username, claims.Username)
		ctx = context.WithValue(ctx, UserID, claims.UserID)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}