package user

import (
	"github.com/QuanCters/backend/internal/model"
	"github.com/QuanCters/backend/internal/service"
	"github.com/QuanCters/backend/pkg/response"
	"github.com/gin-gonic/gin"
)

var UserLoginController = new(cUserLogin)

type cUserLogin struct {
}

func (c *cUserLogin) Login(ctx *gin.Context){
	var params model.LoginInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	codeStatus, output, err := service.UserLogin().Login(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}	
	response.SuccessResponse(ctx, codeStatus, output)
}