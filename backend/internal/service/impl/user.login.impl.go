package impl

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/QuanCters/backend/global"
	"github.com/QuanCters/backend/internal/database"
	"github.com/QuanCters/backend/internal/middlewares"
	"github.com/QuanCters/backend/internal/model"
	"github.com/QuanCters/backend/internal/utils"
	"github.com/QuanCters/backend/internal/utils/auth"
	"github.com/QuanCters/backend/internal/utils/crypto"
	"github.com/QuanCters/backend/pkg/response"
)

type sUserLogin struct {
	r *database.Queries
}


func UserLoginImpl(r *database.Queries) *sUserLogin {
	return &sUserLogin{
		r: r,
	}
}


func (s *sUserLogin) Login(ctx context.Context, in *model.LoginInput) (codeResult int, out model.LoginOutput, err error) {
	// Find User
	userFound, err := s.r.GetUserByEmail(ctx, in.Email)
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}

	// Check Password
	if !crypto.MatchingPassword(userFound.UserPassword, in.Password, userFound.UserSalt) {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("does not match password")
	}

	// Create UUID User
	subToken := utils.GenerateTokenUUID(int(userFound.UserID))

	// Store user info in Redis Hash
	ttl := time.Duration(global.Config.Redis.ExpireTime) * time.Minute
	userKey := fmt.Sprintf("user:%s", subToken)

	// Create pipeline for atomic operations
	pipe := global.Rdb.TxPipeline()
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
		return response.ErrCodeAuthFailed, out, err
	}

	// Create Token
	out.Token, err = auth.CreateToken(subToken, userFound.UserName, userFound.UserID)
	if err != nil {
		return response.ErrCodeAuthFailed, out , err
	}

	return response.ErrCodeSuccess, out, nil
}

func (s *sUserLogin) Logout(ctx context.Context) (codeResult int, out model.LogoutOutput, err error) {
	subToken, ok := ctx.Value(middlewares.SubjectUUIDKey).(string)
	if !ok || subToken == "" {
		// Token already invalid/expired - consider successful logout
		return response.ErrCodeSuccess, out, nil
	}

	userKey := fmt.Sprintf("user:%s", subToken)

	go func (key string) {
		// Create new context with timeout
		delCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		// Perform deletion and handle errors
		if err := global.Rdb.Del(delCtx, key).Err(); err != nil {
			// Use structured logging with more details
			log.Printf("Redis logout deletion failed for key %s: %v", key, err)
			
			// Attempt retry once after short delay
			time.Sleep(100 * time.Millisecond)
			if retryErr := global.Rdb.Del(delCtx, key).Err(); retryErr != nil {
				log.Printf("Retry failed for key %s: %v", key, retryErr)
			}
		} else {
			log.Printf("Successfully logged out token: %s", key)
		}
	}(userKey)

	return response.ErrCodeSuccess, out, nil
}