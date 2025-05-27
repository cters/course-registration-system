package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/QuanCters/backend/global"
	"github.com/QuanCters/backend/internal/database"
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
	log.Println("subtoken:::", subToken)

	// Convert User Info to Json
	userFoundJson, err := json.Marshal(userFound)
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("convert to json failed: %v", err)
	}

	// Add User Info Json to Redis with Token
	err = global.Rdb.Set(ctx, subToken, userFoundJson, time.Duration(1)*time.Minute).Err()
	
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}

	// Create Token
	out.Token, err = auth.CreateToken(subToken)
	if err != nil {
		return response.ErrCodeAuthFailed, out , err
	}

	return response.ErrCodeSuccess, out, nil
}

func (s *sUserLogin) Logout(ctx context.Context, in *model.LogoutInput) (codeResult int, out model.LogoutOutput, err error) {
	return 0, out, nil
}