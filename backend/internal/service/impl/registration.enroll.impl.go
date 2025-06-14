package impl

import (
	"context"
	"encoding/json"
	"time"

	"github.com/QuanCters/backend/global"
	"github.com/QuanCters/backend/internal/database"
	"github.com/QuanCters/backend/internal/model"
	u "github.com/QuanCters/backend/internal/utils/messageQueue"
	"github.com/QuanCters/backend/pkg/response"
	"go.uber.org/zap"
)

type sRegistrationEnroll struct {
	r *database.Queries
}

func RegistrationEnrollImpl(r *database.Queries) *sRegistrationEnroll {
	return &sRegistrationEnroll{
		r: r,
	}
}

func (s *sRegistrationEnroll) Enroll(ctx context.Context, in *model.RegistrationInput) (codeResult int, err error){
	// step 1: call user service to check student exist
	studentFound, err := s.r.CheckStudentExist(ctx, in.StudentID)
	if err != nil || studentFound == 0 {
		global.Logger.Warn("Student not found", zap.Int64("student_id", in.StudentID), zap.Error(err),)
		return response.ErrCodeUserNotFound, err
	}

	// step 2: check for empty slot
	// step 2.1: check for redis

	// step 2.2: if redis doesn't have it, check database
	courseFound, err := s.r.GetCourseById(ctx, in.CourseID)
	if err != nil {
		global.Logger.Warn("Course not found", zap.Int32("course_id", in.CourseID), zap.Error(err),)
		return response.ErrCodeCourseNotFound, err
	}
	emptySlot := courseFound.CourseMaxSlot - courseFound.CourseCurrentSlot
	if emptySlot <= 0 {
		return response.ErrCodeConflict, nil
	}

	// step 3: send message to RabbitMQ
	// step 3.1: create message
	msg := model.RegistrationMessage {
		StudentID: in.StudentID,
		CourseID: in.CourseID,
		Timestamp: time.Now(),
	}
	// step 3.2: convert message to JSON
	msgBody, err := json.Marshal(msg)
	if err != nil {
		global.Logger.Error("Failed to marshal message for RabbitMQ",
			zap.Error(err),
			zap.Int64("student_id", in.StudentID),
			zap.Int32("course_id", in.CourseID),
		)	
		return response.ErrCodeInternal, err
	}
	// step 3.3: send message
	err = u.PublishMessage("registration_exchange", "registration.create", msgBody)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeAccepted, nil
}

func (s *sRegistrationEnroll) Withdraw(ctx context.Context, in *model.RegistrationInput) (codeResult int, err error){
	return 
}