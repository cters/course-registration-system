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
	service.InitUserAdmin(impl.UserAdminImpl(queries))
}