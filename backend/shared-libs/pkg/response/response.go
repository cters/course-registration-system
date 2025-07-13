package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`  // data return null not show
	Error   interface{} `json:"error,omitempty"` // Error return null not show
}

func SuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, APIResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string, err interface{}) {
	c.JSON(code, APIResponse{
		Code:    code,
		Message: message,
		Error:   err,
	})
}


type StatusResponse struct {
	Status  int         `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}


type HandlerFunc func(ctx *gin.Context) (res interface{}, err error)


func Wrap(handler HandlerFunc) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		res, err := handler(ctx)
		if err != nil {
			if apiErr, ok := err.(*APIError); ok {
				ErrorResponse(ctx, apiErr.StatusCode, apiErr.Message, apiErr.Err)
			} else {
				ErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error", err)
			}
			return
		}
		if statusRes, ok := res.(StatusResponse); ok {
			SuccessResponse(ctx, statusRes.Status, statusRes.Message, statusRes.Data)
		} else {
			SuccessResponse(ctx, http.StatusOK, "", res)
		}
	}
}