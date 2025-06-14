package service

import (
	"context"
)

type (
	IRegistrationProcessor interface {
		CreateProcessor(ctx context.Context, msg string) (codeResult int, err error)
		DeleteProcessor(ctx context.Context, message string) (codeResult int, err error)
	}
)

var (
	localRegistrationProcessor IRegistrationProcessor 
)

func RegistrationProcessor() IRegistrationProcessor {
	if localRegistrationProcessor == nil {
		panic("implement localRegistrationProcessor not found for interface IRegistrationProcessor")
	}
	return localRegistrationProcessor
}

func InitRegistrationProcessor(i IRegistrationProcessor) {
	localRegistrationProcessor = i
}
