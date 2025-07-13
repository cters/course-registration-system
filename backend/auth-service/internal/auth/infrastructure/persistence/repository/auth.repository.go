package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/internal/auth/domain/entity"
	"gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/internal/auth/domain/repository"
	"gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/internal/auth/infrastructure/persistence/sqlc"
)

type authRepository struct {
	db *sql.DB
	queries *sqlc.Queries
}

// CreateAccountWithTx implements repository.AuthRepository.
func (ar *authRepository) CreateAccountWithTx(ctx context.Context, account *entity.Account) (int32, error) {
	tx, err := ar.db.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	qtx := ar.queries.WithTx(tx)
	var UserID int32
	defer func() {
		if p := recover(); p != nil {
			// Xử lý panic: rollback và re-panic
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			// Xử lý lỗi: rollback
			_ = tx.Rollback()
		} else {
			// Không có lỗi: commit
			err = tx.Commit()
			if err != nil {
				// Xử lý lỗi khi commit
				err = fmt.Errorf("failed to commit transaction: %w", err)
			}
		}
	}()

	params := sqlc.AddUserParams{
		UserEmail:    account.UserEmail,
		UserAccount: account.UserAccount,
		UserPassword: account.UserPassword,
		UserSalt:   account.UserSalt,
		UserName: account.UserName,
		UserPhone:    account.UserPhone,
	}

	UserID, err = qtx.AddUser(ctx, params)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to create user: %w", err)
	}

	return UserID, nil
}

func (ar *authRepository) CheckUserExist(ctx context.Context, email string) (bool, error) {
	count, err := ar.queries.CheckUserExist(ctx, email)
	if err != nil {
		return false, fmt.Errorf("failed to check user existence: %w", err)
	}
	return count > 0, nil
}

func (ar *authRepository) GetUserByEmail(ctx context.Context, email string) (*entity.Account, error) {
	account, err := ar.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	if account.UserID == 0 {
		return nil, nil // User not found
	}

	return &entity.Account{
		UserID:       account.UserID,
		UserEmail:    account.UserEmail,
		UserAccount:  account.UserAccount,
		UserPassword: account.UserPassword,
		UserSalt:     account.UserSalt,
		UserName:     account.UserName,
		UserPhone:    account.UserPhone,
	}, nil
}

func NewAuthRepository(db *sql.DB) repository.AuthRepository {
	return &authRepository{
		queries: sqlc.New(db),
		db: db,
	}
}
