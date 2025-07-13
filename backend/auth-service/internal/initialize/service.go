package initialize

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
	"gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/internal/auth/application/service"
	"gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/internal/auth/controller/http"
	authRepo "gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/internal/auth/infrastructure/persistence/repository"
	setting "gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/pkg"
	"gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/utils/auth"
	"go.uber.org/zap"
)

// initializes service, repository, and handler for auth
func InitAuth(
	config setting.Config,
	db *sql.DB, 
	logger *zap.Logger,
	redisClient *redis.Client,
) *http.AuthHandler {
	cookieManager := auth.NewCookieManager(config.JwtExpiration)
	tokenManager := auth.NewTokenManager(&auth.JWTCfg{
		SecretKey: config.SecretKey,
		TokenExpiration: config.JwtExpiration,
	})
	authRepo := authRepo.NewAuthRepository(db)
	service := service.NewAuthService(
		authRepo, 
		config,
		tokenManager,
		logger,
		redisClient,	
	)
	handler := http.NewAuthHandler(service, cookieManager)
	return handler
}