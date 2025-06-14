package impl

import (
	"context"
	"fmt"

	"github.com/QuanCters/backend/internal/database"
	"github.com/QuanCters/backend/internal/model"
	"github.com/QuanCters/backend/pkg/response"
)

type sCourseAdmin struct {
	r *database.Queries
}

func CourseAdminImpl(r *database.Queries) *sCourseAdmin {
	return &sCourseAdmin{
		r: r,
	}
}

func (s *sCourseAdmin) CreateCourse(ctx context.Context, in *model.CourseInput) (codeResult int,  err error){
	subjectFound, err := s.r.CheckSubjectExist(ctx, in.SubjectID)
	if err != nil || subjectFound == 0 {
		return response.ErrCodeSubjectNotFound, err
	}

	params := database.AddCourseParams{
		SubjectID: in.SubjectID,
		CourseTerm: in.CourseTerm,
		CourseMaxSlot: in.CourseMaxSlot,
	}

	_, err = s.r.AddCourse(ctx, params)
	if err != nil {
		return response.ErrCodeInternal, fmt.Errorf("failed to create course: %w", err)
	}

	return response.ErrCodeSuccess, nil
}

func (s *sCourseAdmin) UpdateCourse(ctx context.Context, in *model.Course) (codeResult int, err error){
	courseFound, err := s.r.CheckCourseExist(ctx, in.CourseID);
	if err != nil || courseFound == 0 {
		return response.ErrCodeCourseNotFound, err
	}

	params := database.UpdateCourseByIdParams{
		CourseID: in.CourseID,
		SubjectID: in.SubjectID,
		CourseTerm: in.CourseTerm,
		CourseMaxSlot: in.CourseMaxSlot,
		CourseCurrentSlot: in.CourseCurrentSlot,
	}

	_, err = s.r.UpdateCourseById(ctx, params)
	if err != nil {
		return response.ErrCodeInternal, fmt.Errorf("failed to update course: %w", err)
	}
	return response.ErrCodeSuccess, nil
}

func (s *sCourseAdmin) DeleteCourse(ctx context.Context, courseId int) (codeResult int, err error){
	courseFound, err := s.r.CheckCourseExist(ctx, int32(courseId));
	if err != nil || courseFound == 0 {
		return response.ErrCodeCourseNotFound, err
	}

	_, err = s.r.DeleteCourseById(ctx, int32(courseId))	
	if err != nil {
		return response.ErrCodeInternal, fmt.Errorf("failed to delete course: %w", err)
	}
	return response.ErrCodeSuccess, nil
}