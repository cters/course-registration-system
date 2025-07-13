package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	httpDto "gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/internal/auth/controller/dto"
	"gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/internal/auth/domain/entity"
	authRepo "gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/internal/auth/domain/repository"
	setting "gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/pkg"
	"gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/utils"
	"gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/utils/auth"
	crypto "gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/utils/crypto"
	"go.uber.org/zap"
)

type authService struct {
	authRepo authRepo.AuthRepository	
	config setting.Config
	tokenManager *auth.TokenManager
	logger *zap.Logger
	redisClient *redis.Client
}

// CreateAccount implements AuthService.
func (a *authService) CreateAccount(ctx context.Context, account *httpDto.CreateAccountReq) (int, error) {
	// Check username exists 
	userFound, err := a.authRepo.CheckUserExist(ctx, account.Email)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if userFound {
		return http.StatusConflict, errors.New("email already exist")
	}

	// role := "guest"
	// if account.Role != nil {
	// 	role = *account.Role
	// }

	salt, err := crypto.GenerateSalt(16)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	hashedPassword := crypto.HashPassword(account.Password, salt)

	newAccount := &entity.Account{
		UserEmail: account.Email,
		UserAccount: strings.Split(account.Email, "@")[0],
		UserPassword: hashedPassword,
		UserSalt: salt,
		UserName: account.Name,
		UserPhone: account.Phone,
	}

	_, err = a.authRepo.CreateAccountWithTx(ctx, newAccount)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to create account: %w", err) 
	}

	// _, err = qtx.AssignRoleToUser(ctx, database.AssignRoleToUserParams{
	// 	UserID: _UserID,
	// 	RoleID: roleFound.RoleID,
	// })
	// if err != nil {
	// 	return  response.ErrCodeInternal, fmt.Errorf("failed to assign role: %w",err)
	// }
	return http.StatusCreated, nil
}

// Login implements AuthService.
func (a *authService) Login(ctx context.Context, login *httpDto.UserLoginReq) (int, *httpDto.UserLoginRes, error) {
	userFound, err := a.authRepo.GetUserByEmail(ctx, login.Email)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if !crypto.MatchingPassword(userFound.UserPassword, login.Password, userFound.UserSalt) {
		return http.StatusUnauthorized, nil, errors.New("incorrect password")
	}

	subToken := utils.GenerateTokenUUID(int(userFound.UserID))
	ttl := time.Duration(a.config.RedisExpire)*time.Minute
	userKey := fmt.Sprintf("user:%s", subToken)

	// Create pipeline for atomic operations
	pipe := a.redisClient.TxPipeline()
	pipe.HSet(ctx, userKey, map[string]interface{}{
		"user_id": userFound.UserID,
		"user_account": userFound.UserAccount,
		"user_name": userFound.UserName,
		"user_email": userFound.UserEmail,
		"user_phone": userFound.UserPhone,
	})
	pipe.Expire(ctx, userKey, ttl)

	// Execute pipeline commands atomically
	if _, err := pipe.Exec(ctx); err != nil {
		log.Printf("failed to save user data to redis: %v", err)
		return http.StatusUnauthorized, &httpDto.UserLoginRes{}, err
	}

	// Create Token
	token, err := a.tokenManager.CreateToken(subToken, userFound.UserName, userFound.UserID)
	if err != nil {
		return http.StatusUnauthorized, &httpDto.UserLoginRes{} , err
	}

	return http.StatusOK, &httpDto.UserLoginRes{
		UserID: userFound.UserID,
		UserAccount: userFound.UserAccount,
		Username: userFound.UserName,
		UserEmail: userFound.UserEmail,
		UserPhone: userFound.UserPhone,
		Token: token,
	}, nil
}

// Logout implements AuthService.
func (a *authService) Logout(ctx context.Context, token string) error {
	return nil
}

func NewAuthService(
	authRepo authRepo.AuthRepository, 
	config setting.Config,
	tokenManager *auth.TokenManager,
	logger *zap.Logger,
	redisClient *redis.Client,
) AuthService {
	return &authService{
		authRepo: authRepo,
		config: config,
		tokenManager: tokenManager,
		logger: logger,
		redisClient: redisClient,
	}
}
