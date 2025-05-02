package initialize

import (
	"github.com/QuanCters/backend/global"
	"github.com/QuanCters/backend/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}