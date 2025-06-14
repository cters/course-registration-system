package impl

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/QuanCters/backend/global"
	"github.com/QuanCters/backend/internal/database"
	"github.com/QuanCters/backend/internal/middlewares"
	"github.com/QuanCters/backend/internal/model"
	"github.com/QuanCters/backend/pkg/response"
)

type sUserInfo struct {
	r *database.Queries
}

func UserInfoImpl(r *database.Queries) *sUserInfo {
	return &sUserInfo{
		r: r,
	}
}

func (s *sUserInfo) GetMyInfo(ctx context.Context) (codeResult int, out model.UserOutput, err error)  {
	// Get from redis first
	subToken, ok := ctx.Value(middlewares.SubjectUUIDKey).(string)
	if !ok || subToken == "" {
		return response.ErrCodeInternal, model.UserOutput{}, fmt.Errorf("missing subject UUID in context")
	}
	userKey := fmt.Sprintf("user:%s", subToken)

	result, err := global.Rdb.HGetAll(ctx, userKey).Result()
	if err == nil && len(result) > 0 {
		userID, _ := strconv.ParseInt(result["user_id"], 10, 64)
		return response.ErrCodeSuccess, model.UserOutput{
			UserID: int32(userID),
			UserAccount: result["user_account"],
			UserEmail:   result["user_email"],
			UserName:    result["user_name"],
			UserPhone:   result["user_phone"],
		}, nil
	}
	
	// Not found in redis
	user_id, ok := ctx.Value(middlewares.UserID).(int32)
	if !ok {
		return response.ErrCodeInternal, model.UserOutput{}, fmt.Errorf("invalid user ID in context")
	}

	dbUser, err := s.r.GetUserById(ctx, user_id)
	if err != nil {
		return response.ErrCodeInternal, model.UserOutput{}, fmt.Errorf("failed to get user information for user_id %d: %w", user_id, err)
	}

	// Convert to output model
	out = model.UserOutput{
		UserID: dbUser.UserID,
		UserAccount: dbUser.UserAccount,
		UserEmail: dbUser.UserEmail,
		UserName:  dbUser.UserName,
		UserPhone: dbUser.UserPhone,
	}

	// update Redis cache (async to not block response)
	go func(u model.UserOutput) {
		bgCtx := context.Background()
		ttl := time.Duration(global.Config.Redis.ExpireTime) * time.Minute

		userFields := map[string]interface{}{
			"user_id":      strconv.FormatInt(int64(u.UserID), 10),
			"user_account": u.UserAccount,
			"user_email":   u.UserEmail,
			"user_name":    u.UserName,
			"user_phone":   u.UserPhone,
		}

		pipe := global.Rdb.TxPipeline()
		pipe.HSet(bgCtx, userKey, userFields)
		pipe.Expire(bgCtx, userKey, ttl)
		
		if _, err := pipe.Exec(bgCtx); err != nil {
			log.Printf("Redis cache update failed: %v", err)
		}
		
	}(out)

	return response.ErrCodeSuccess, out, nil
}

