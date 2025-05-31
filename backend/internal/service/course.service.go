package service

import (
	"context"

	"github.com/QuanCters/backend/internal/model"
)

type (
	ICourseAdmin interface {
		CreateCourse(ctx context.Context, in *model.CourseInput) (codeResult int,  err error)
		UpdateCourse(ctx context.Context, in *model.CourseInput) (codeResult int, err error)
		DeleteCourse(ctx context.Context, courseId int) (codeResult int, err error)
	}

	ICourseInfo interface {
		GetCourseById(ctx context.Context, courseId int) (codeResult int, out *model.Course, err error)
		GetCourseBySubjectIdAndTerm(ctx context.Context, subjectId int, term string) (codeResult int, out *model.Course, err error)
	}
	
	ICourseRegistration interface {
		CheckRequirement(ctx context.Context, studentId int, courseId int) (codeResult int, meets bool, err error)
	}
)

var (
	localCourseAdmin ICourseAdmin
	localCourseInfo  ICourseInfo
	localCourseEnroll ICourseRegistration
)

func CourseAdmin() ICourseAdmin {
	if localCourseAdmin == nil {
		panic("implement localCourseAdmin not found for interface ICourseAdmin")
	}
	return localCourseAdmin
}

func InitCourseAdmin(i ICourseAdmin) {
	localCourseAdmin = i
}

func CourseInfo() ICourseInfo {
	if localCourseInfo == nil {
		panic("implement localCourseInfo not found for interface ICourseInfo")
	}
	return localCourseInfo
}

func InitCourseInfo(i ICourseInfo) {
	localCourseInfo = i
}

func CourseEnroll() ICourseRegistration {
	if localCourseEnroll == nil {
		panic("implement localCourseEnroll not found for interface ICourseEnroll")
	}
	return localCourseEnroll
}

func InitCourseEnroll(i ICourseRegistration) {
	localCourseEnroll = i
}