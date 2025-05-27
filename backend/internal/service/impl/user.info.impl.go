package impl

import (
	"context"

	"github.com/QuanCters/backend/internal/database"
)

type sUserInfo struct {
	r *database.Queries
}

func UserInfoImpl(r *database.Queries) *sUserInfo {
	return &sUserInfo{
		r: r,
	}
}

func (s *sUserInfo) GetInfoByUserId(ctx context.Context) error {
	return nil
}

