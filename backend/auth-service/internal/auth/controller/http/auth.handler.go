package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/internal/auth/application/service"
	httpDto "gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/internal/auth/controller/dto"
	"gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/utils"
	"gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/utils/auth"
	"gitlab.com/dacn9315724/course-registration-ddd/backend/shared-libs/pkg/response"
)

type AuthHandler struct {
	service service.AuthService
	cookieManager *auth.CookieManager
}


func (ah *AuthHandler) CreateAccount(ctx *gin.Context) (res interface{}, err error) {
	// Implementation for creating an account
	var req httpDto.CreateAccountReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, response.NewAPIError(http.StatusBadRequest, "Invalid request data", err.Error()) 
	}
	validation, exists := ctx.Get("validation")
	if !exists {
		return nil, response.NewAPIError(http.StatusBadRequest, "Invalid request", "Validation not found")
	}
	if apiErr := utils.ValidateStruct(req, validation.(*validator.Validate)); apiErr != nil {
		return nil, apiErr
	}

	userID, err := ah.service.CreateAccount(ctx, &req)
	if err != nil {
		return nil, response.NewAPIError(http.StatusInternalServerError, "Failed to create account", err.Error())
	}
	
	return response.StatusResponse{
		Status: http.StatusCreated,
		Message: "Account created successfully",
		Data: &httpDto.CreateAccountRes{UserID: userID},
	}, nil
}


func (ah *AuthHandler) Login(ctx *gin.Context) (res interface{}, err error) {
	// Implementation for user login
	var req httpDto.UserLoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, response.NewAPIError(http.StatusBadRequest, "Invalid request data", err.Error())
	}

	validation, exists := ctx.Get("validation")
	if !exists {
		return nil, response.NewAPIError(http.StatusBadRequest, "Invalid request", "Validation not found")
	}

	if apiErr := utils.ValidateStruct(req, validation.(*validator.Validate)); apiErr != nil {
		return nil, apiErr
	}

	statusCode, loginRes, err := ah.service.Login(ctx, &req)
	if err != nil {
		return nil, response.NewAPIError(http.StatusInternalServerError, "Failed to login", err.Error())
	}	

	ah.cookieManager.SetTokenCookie(ctx, loginRes.Token)

	return response.StatusResponse{
		Status: statusCode,
		Message: "Login successful",
		Data: loginRes,
	}, nil
}


func (ah *AuthHandler) Logout(ctx *gin.Context) (res interface{}, err error) {
	// Implementation for user logout
	token, err := ah.cookieManager.GetTokenFromCookie(ctx)
	if err != nil {
		return nil, response.NewAPIError(http.StatusUnauthorized, "Unauthorized", "Invalid or missing token")
	}

	if err := ah.service.Logout(ctx, token); err != nil {
		return nil, response.NewAPIError(http.StatusInternalServerError, "Failed to logout", err.Error())
	}

	ah.cookieManager.ClearTokenCookie(ctx)

	return response.StatusResponse{
		Status: http.StatusOK,
		Message: "Logout successful",
	}, nil
}


func NewAuthHandler(service service.AuthService, cookieManager *auth.CookieManager) *AuthHandler {
	return &AuthHandler{
		service: service,
		cookieManager: cookieManager,
	}
}