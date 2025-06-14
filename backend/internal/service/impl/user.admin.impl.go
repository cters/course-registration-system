package impl

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/QuanCters/backend/internal/database"
	"github.com/QuanCters/backend/internal/model"
	"github.com/QuanCters/backend/internal/utils/crypto"
	"github.com/QuanCters/backend/pkg/response"
)

type sUserAdmin struct {
	db *sql.DB
	r *database.Queries
}

func UserAdminImpl(db *sql.DB) *sUserAdmin{
	return &sUserAdmin{
		db: db,
		r: database.New(db),
	}
}

func (s *sUserAdmin) Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error) {
	// 1. Kiểm tra người dùng đã tồn tại chưa
	userFound, err := s.r.CheckUserExist(ctx, in.Email)
	if err != nil {
		return response.ErrCodeUserHasExists, err
	}
	if userFound > 0 {
		return response.ErrCodeUserHasExists, fmt.Errorf("user has already registered")
	}

	// 2. Xác định và kiểm tra vai trò hợp lệ
	role := "guest"
	if in.Role != "" {
		role = in.Role
	}

	roleNullString := sql.NullString{
		String: role,
		Valid: role != "",
	}
	roleFound, err := s.r.GetRoleByName(ctx, roleNullString)
	if err != nil {
		if err == sql.ErrNoRows{
			return response.ErrCodeUnprocessableEntity, fmt.Errorf("role '%s' does not exist", role)
		}
		return response.ErrCodeInternal, fmt.Errorf("failed to retrieve role information: %w", err)
	}

	// 3. Tạo salt và hash mật khẩu
	salt, err := crypto.GenerateSalt(16)
	if err != nil {
		return response.ErrCodeInternal, fmt.Errorf("failed to generate salt: %w", err)
	}
	hashedPassword := crypto.HashPassword(in.Password, salt)

	// 4. Bắt đầu ACID transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return response.ErrCodeInternal, fmt.Errorf("failed to begin transaction: %w", err)
	}

	// 5. Tạo một đối tượng Queries mới sử dụng giao dịch này
	qtx := s.r.WithTx(tx)
	// Sử dụng defer để đảm bảo giao dịch được kết thúc
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
				codeResult = response.ErrCodeInternal
				err = fmt.Errorf("failed to commit transaction: %w", err)
			}
		}
	}()

	// 6. Chuẩn bị tham số để thêm người dùng
	params := database.AddUserParams{
		UserAccount:  strings.Split(in.Email, "@")[0],
		UserEmail:    in.Email,
		UserPassword: hashedPassword,
		UserSalt:     salt,
		UserName: 	  in.Name,
		UserPhone:    in.Phone,
	}

	// 7. Thêm người dùng vào cơ sở dữ liệu
	_UserID, err := qtx.AddUser(ctx, params)
	if err != nil {
		return response.ErrCodeInternal, fmt.Errorf("failed to create user: %w", err)
	}

	// 8. Gán vai trò cho người dùng
	_, err = qtx.AssignRoleToUser(ctx, database.AssignRoleToUserParams{
		UserID: _UserID,
		RoleID: roleFound.RoleID,
	})
	if err != nil {
		return  response.ErrCodeInternal, fmt.Errorf("failed to assign role: %w",err)
	}

	return response.ErrCodeCreated, nil
}