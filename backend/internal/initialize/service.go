package initialize

import (
	"github.com/QuanCters/backend/global"
	"github.com/QuanCters/backend/internal/database"
	"github.com/QuanCters/backend/internal/service"
	"github.com/QuanCters/backend/internal/service/impl"
)

func InitServiceInterface() {
	queries := database.New(global.Pdbc)
	// User Service Interface
	service.InitUserLogin(impl.UserLoginImpl(queries))
	service.InitUserAdmin(impl.UserAdminImpl(global.Pdbc))
	service.InitUserInfo(impl.UserInfoImpl(queries))
	// Course Service Interface 
	service.InitCourseAdmin(impl.CourseAdminImpl(queries))
	service.InitCourseInfo(impl.CourseInfoImpl(queries))
	// Registration Service Interface
	service.InitRegistrationEnroll(impl.RegistrationEnrollImpl(queries))
	service.InitRegistrationProcessor(impl.RegistrationProcessorImpl(global.Pdbc))
}