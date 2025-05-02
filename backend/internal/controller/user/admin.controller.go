package user

import (
	"github.com/QuanCters/backend/internal/model"
	"github.com/QuanCters/backend/internal/service"
	"github.com/QuanCters/backend/pkg/response"
	"github.com/gin-gonic/gin"
)

var UserAdminController = new(cUserAdmin)

type cUserAdmin struct {
}

func (c *cUserAdmin) Register(ctx *gin.Context) {
	var params model.RegisterInput
	if err := ctx.ShouldBindJSON(&params);err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	codeStatus, err := service.UserAdmin().Register(ctx, &params)

	if err != nil {
		response.ErrorResponse(ctx, codeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeStatus, nil)	
}