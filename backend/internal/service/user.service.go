package service

import (
	"context"

	"github.com/QuanCters/backend/internal/model"
)

type (
	IUserLogin interface {
		Login(ctx context.Context, in *model.LoginInput) (codeResult int, out model.LoginOutput, err error) 
		Logout(ctx context.Context, in *model.LogoutInput) (codeResult int, out model.LogoutOutput, err error)
	}
	
	IUserInfo interface {
		GetInfoByUserId(ctx context.Context) error
	}
	
	IUserAdmin interface {
		Register(ctx context.Context, in *model.RegisterInput) (codeResult int,  err error)
	}
)

var (
	localUserAdmin IUserAdmin
	localUserInfo IUserInfo
	localUserLogin IUserLogin
)

func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		panic("Implement localUserAdmin not found for interface IUserAdmin")
	}

	return localUserAdmin
}

func InitUserAdmin(i IUserAdmin) {
	localUserAdmin = i
}

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("implement localUserInfo not found for interface IUserInfo")
	}
	return localUserInfo
}

func InitUserInfo(i IUserInfo) {
	localUserInfo = i
}

func UserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("implement localUserLogin not found for interface IUserLogin")
	}
	return localUserLogin
}

func InitUserLogin(i IUserLogin) {
	localUserLogin = i
}
