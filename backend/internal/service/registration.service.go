package service

import (
	"context"

	"github.com/QuanCters/backend/internal/model"
)

type (
	IRegistrationEnroll interface {
		Enroll(ctx context.Context, in *model.RegistrationInput) (codeResult int, err error)
		Withdraw(ctx context.Context, in *model.RegistrationInput) (codeResult int, err error)
	}
)

var (
	localRegistrationEnroll IRegistrationEnroll
)

func RegistrationEnroll() IRegistrationEnroll {
	if localRegistrationEnroll == nil {
		panic("implement localRegistrationEnroll not found for interface IRegistrationEnroll")
	}
	return localRegistrationEnroll
}

func InitRegistrationEnroll(i IRegistrationEnroll) {
	localRegistrationEnroll = i
}