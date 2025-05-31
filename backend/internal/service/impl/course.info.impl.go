package impl

import (
	"context"

	"github.com/QuanCters/backend/internal/database"
	"github.com/QuanCters/backend/internal/model"
	"github.com/QuanCters/backend/pkg/response"
)

type sCourseInfo struct {
	r *database.Queries
}

func  CourseInfoImpl(r *database.Queries) *sCourseInfo {
	return &sCourseInfo{
		r: r,
	}
}

func (s *sCourseInfo) GetCourseById(ctx context.Context, courseId int) (codeResult int, out *model.Course, err error){
	course, err := s.r.GetCourseById(ctx, int64(courseId))
	if err != nil {
		return response.ErrCodeCourseNotFound, nil, err
	}
	result := &model.Course{
		CourseID: course.CourseID,
		SubjectID: course.SubjectID,
		CourseTerm: course.CourseTerm,
		CourseMaxSlot: course.CourseMaxSlot,
		CourseCurrentSlot: course.CourseCurrentSlot,
	}
	return response.ErrCodeSuccess, result, nil 
}
	
func (s *sCourseInfo) GetCourseBySubjectIdAndTerm(ctx context.Context, subjectID string, term string) (codeResult int, out *model.Course, err error){
	params := database.GetCourseBySubjectIdAndTermParams{
		SubjectID: subjectID,
		CourseTerm: term,
	}

	course, err := s.r.GetCourseBySubjectIdAndTerm(ctx, params)
	if err != nil {
		return response.ErrCodeCourseNotFound, nil, err
	}

	result := &model.Course{
		CourseID: course.CourseID,
		SubjectID: course.SubjectID,
		CourseTerm: course.CourseTerm,
		CourseMaxSlot: course.CourseMaxSlot,
		CourseCurrentSlot: course.CourseCurrentSlot,
	}
	return response.ErrCodeSuccess, result, nil 
}