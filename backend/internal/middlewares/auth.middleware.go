package middlewares

import (
	"context"
	"log"

	"github.com/QuanCters/backend/internal/utils/auth"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.URL.Path
		log.Println("uri request: ", uri)
		jwtToken, valid := auth.ExtractBearerToken(c)
		if !valid {
			c.AbortWithStatusJSON(401, gin.H{"code": 401, "err": "Unauthorized", "description": ""})
			return
		}

		claims, err := auth.VerifyTokenSubject(jwtToken)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"code": 401, "err": "Invalid Token", "description": ""})
		}

		log.Println("claims::: UUID::", claims.Subject)
		ctx := context.WithValue(c.Request.Context(), "subjectUUID", claims.Subject)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}