package user

import (
	"github.com/QuanCters/backend/internal/service"
	"github.com/QuanCters/backend/pkg/response"
	"github.com/gin-gonic/gin"
)

var UserInfoController = new(cUserInfo)

type cUserInfo struct {
}

func (c *cUserInfo) GetMyInfo(ctx *gin.Context) {
	codeStatus, output, err := service.UserInfo().GetMyInfo(ctx.Request.Context())
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeStatus, output)
}