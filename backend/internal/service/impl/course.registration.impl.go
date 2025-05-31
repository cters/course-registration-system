package impl

import (
	"context"

	"github.com/QuanCters/backend/internal/database"
)

type sCourseRegistration struct {
	r *database.Queries
}

func CourseRegistrationImpl(r *database.Queries) *sCourseRegistration {
	return &sCourseRegistration{
		r: r,
	}
}

func (s *sCourseRegistration) CheckRequirement(ctx context.Context, studentId int, courseId int) (codeResult int, meets bool, err error){
	return
}