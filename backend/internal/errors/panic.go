package errors

import (
	"github.com/QuanCters/backend/global"
	"github.com/QuanCters/backend/pkg/logger"
	"go.uber.org/zap"
)

func MustCheck(err error, msg string) {
	if err != nil {
		// Logger được inject hoặc khởi tạo trong package này
		global.Logger.Error(msg, zap.Error(err))
		panic(err)
	}
}

func Must(logger *logger.LoggerZap, err error, component string) {
	if err != nil {
		panic(err)
	}
}