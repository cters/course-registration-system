package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/QuanCters/backend/global"
	"github.com/QuanCters/backend/internal/database"
	"github.com/QuanCters/backend/internal/model"
	"github.com/QuanCters/backend/pkg/response"
	"go.uber.org/zap"
)

type sRegistrationProcessor struct {
	db *sql.DB
	r *database.Queries
}

func RegistrationProcessorImpl(db *sql.DB) *sRegistrationProcessor {
	return &sRegistrationProcessor{
		db: db,
		r: database.New(db),
	}
}

func (s *sRegistrationProcessor) CreateProcessor(ctx context.Context, msg string) (codeResult int, err error) {
	// 1. parse the message to extract registration details
	var registrationMessage model.RegistrationMessage
	err = json.Unmarshal([]byte(msg), &registrationMessage)
	if err != nil {
		global.Logger.Error("Failed to unmarshal message", zap.Error(err))
		return response.ErrCodeInternal, err
	}

	// 2. Begin ACID transaction
	tx, err := s.db.BeginTx(ctx, nil)	
	if err != nil {
		global.Logger.Error("Failed to begin transaction", zap.Error(err))
		return response.ErrCodeInternal, err
	}
	qtx := s.r.WithTx(tx)
	 
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
			if err != nil {
				codeResult = response.ErrCodeInternal
				err = fmt.Errorf("failed to commit transaction: %w", err)
			}
		}
	}()

	// 3. Insert new registration record 
	params := database.CreateRegistrationParams{
		StudentID: registrationMessage.StudentID,
		CourseID: int32(registrationMessage.CourseID),
	}
	_, err = qtx.CreateRegistration(ctx, params)
	if err != nil {
		return response.ErrCodeInternal, fmt.Errorf("failed to register: %w", err)
	}

	return response.ErrCodeSuccess, nil
}

func (s *sRegistrationProcessor) DeleteProcessor(ctx context.Context, message string) (codeResult int, err error) {
	return
}