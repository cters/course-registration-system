package repo

import (
	"github.com/QuanCters/backend/global"
	"github.com/QuanCters/backend/internal/database"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
	sqlc *database.Queries
}

func (ur *userRepository) GetUserByEmail(email string) bool {
	return false
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.Pdbc),
	}
}