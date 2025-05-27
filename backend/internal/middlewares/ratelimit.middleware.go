package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/QuanCters/backend/global"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	redisStore "github.com/ulule/limiter/v3/drivers/store/redis"
)


type RateLimiter struct {
	globalRateLimiter *limiter.Limiter
	publicAPIRateLimiter *limiter.Limiter
	privateAPIRateLimiter *limiter.Limiter
}

func NewRateLimiter() *RateLimiter {
	rateLimiter := &RateLimiter{
		globalRateLimiter: rateLimiter("5000-S"),
		publicAPIRateLimiter: rateLimiter("2000-S"),
		privateAPIRateLimiter: rateLimiter("1000-S"),
	}

	return rateLimiter
}

func rateLimiter(interval string) *limiter.Limiter{
	store, err := redisStore.NewStoreWithOptions(global.Rdb, limiter.StoreOptions{
		Prefix: "rate-limiter",
		MaxRetry: 3,
		CleanUpInterval: time.Hour,
	})
	if err != nil {
		return nil
	}
	rate, err := limiter.NewRateFromFormatted(interval)
	if err != nil {
		panic (err)
	}
	instance := limiter.New(store, rate)
	return instance
}

// GLOBAL LIMITER
func (rl *RateLimiter) GlobalRateLimiter() gin.HandlerFunc{
	return func(c *gin.Context) {
		key := "global"
		limitContext, err := rl.globalRateLimiter.Get(c, key)
		if err != nil {
			fmt.Println("Failed to check rate limit GLOBAL", err)
			c.Next()
			return
		}

		if limitContext.Reached {
			log.Printf("Rate limit reached GOAL %s", key)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit reached GLOBAL, try later"})
			return 
		}
		c.Next()
	}
}

// PUBLIC API LIMITER
func (rl *RateLimiter) PublicAPIRateLimiter() gin.HandlerFunc{
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		rateLimiterPath := rl.filterLimitUrlPath(urlPath)
		if rateLimiterPath != nil {
			key := fmt.Sprintf("%s-%s", "111-222-333-444", urlPath)	
			limitContext, err := rateLimiterPath.Get(c, key)
			if err != nil {
				fmt.Println("Failed to check rate limit public", err)
				c.Next()
				return
			}
			if limitContext.Reached {
				log.Printf("Rate limit reached GOAL %s", key)
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit reached GLOBAL, try later"})
				return 
			}
		}
		c.Next()
	}
}

// PRIVATE API LIMITER
func (rl *RateLimiter) PrivateAPIRateLimiter() gin.HandlerFunc{
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		rateLimiterPath := rl.filterLimitUrlPath(urlPath)
		if rateLimiterPath != nil {
			userId := 1001
			key := fmt.Sprintf("%d-%s", userId, urlPath)	
			limitContext, err := rateLimiterPath.Get(c, key)
			if err != nil {
				fmt.Println("Failed to check rate limit public", err)
				c.Next()
				return
			}
			if limitContext.Reached {
				log.Printf("Rate limit reached GOAL %s", key)
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit reached GLOBAL, try later"})
				return 
			}
		}
		c.Next()
	}
}

// Filter path
func (rl *RateLimiter) filterLimitUrlPath(urlPath string) *limiter.Limiter {
	if urlPath == "/api/v1/user/login" || urlPath == "/ping/80" {
		return rl.publicAPIRateLimiter
	} else if urlPath == "/api/v1/user/register" || urlPath == "/ping/50"{
		return rl.privateAPIRateLimiter
	} else {
		return rl.globalRateLimiter
	}
}