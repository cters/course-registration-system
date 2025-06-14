package middlewares

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"slices"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type ErrorResponse struct {
	Error string `json:"error"`
	Message string `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}


func init() {
	// Khởi tạo một validator mới
	validate = validator.New()
	
	/* 
		thêm vào validator cách đọc tên trường từ JSON tag
		nếu struct định nghĩa `json:"first_name"` 
		thì sẽ báo lỗi "first_name" thay vì FirstName
	*/
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	
	// thêm 1 vài luật validate khác
	registerCustomValidators()
}

// Register custom validation rules
func registerCustomValidators() {
	// Custom validator for role names
	validate.RegisterValidation("valid_role", func(fl validator.FieldLevel) bool {
		validRoles := []string{"admin", "user", "operator", "guest", "instructor"}
		role := fl.Field().String()
		return slices.Contains(validRoles, role)
	})

	// Custom validator for strong password
	validate.RegisterValidation("strong_password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		return len(password) > 4;	
	})
}

func ValidatorMiddleware(structType any) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a new instance of the struct
		structValue := reflect.New(reflect.TypeOf(structType)).Interface()
		
		// Bind JSON to struct
		if err := c.ShouldBindJSON(structValue); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Error:   "invalid_request",
				Message: "Invalid JSON format",
				Details: map[string]string{"binding_error": err.Error()},
			})
			c.Abort()
			return
		}
		
		// Validate the struct
		if err := validate.Struct(structValue); err != nil {
			validationErrors := make(map[string]string)
			
			for _, err := range err.(validator.ValidationErrors) {
				field := err.Field()
				tag := err.Tag()
				param := err.Param()
				
				validationErrors[field] = getErrorMessage(field, tag, param)
			}
			
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Error:   "validation_failed",
				Message: "Request validation failed",
				Details: validationErrors,
			})
			c.Abort()
			return
		}
		
		// Store validated data in context
		c.Set("validated_data", structValue)
		c.Next()
	}
}

// Tạo ra một middleware để kiểm tra dữ liệu JSON
// T là một kiểu dữ liệu bất kỳ (ví dụ: User, Product,...)
func ValidateJSON[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data T
		
		// 1. Đọc JSON từ request và gán vào biến data
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Error:   "invalid_request",
				Message: "Invalid JSON format",
				Details: map[string]string{"binding_error": err.Error()},
			})
			c.Abort()
			return
		}
		
		// 2. Dùng validator để kiểm tra struct `data`
		if err := validate.Struct(&data); err != nil {
			validationErrors := make(map[string]string)
			
			for _, err := range err.(validator.ValidationErrors) {
				field := err.Field()
				tag := err.Tag()
				param := err.Param()
				
				validationErrors[field] = getErrorMessage(field, tag, param)
			}
			
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Error:   "validation_failed",
				Message: "Request validation failed",
				Details: validationErrors,
			})
			c.Abort()
			return
		}
		
		//3. Nếu ok, lưu dữ liệu vào context
		c.Set("validated_data", data)
		c.Next()
	}
}

// Tạo ra một middleware kiểm tra tham số trên URL 
func ValidateQuery(rules map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		errors := make(map[string]string)
		
		for param, rule := range rules {
			value := c.Query(param)
			
			if err := validate.Var(value, rule); err != nil {
				for _, err := range err.(validator.ValidationErrors) {
					errors[param] = getErrorMessage(param, err.Tag(), err.Param())
				}
			}
		}
		
		if len(errors) > 0 {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Error:   "validation_failed",
				Message: "Query parameter validation failed",
				Details: errors,
			})
			c.Abort()
			return
		}
		
		c.Next()
	}
}

// Helper function to generate user-friendly error messages
func getErrorMessage(field, tag, param string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", field, param)
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long", field, param)
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters long", field, param)
	case "numeric":
		return fmt.Sprintf("%s must be numeric", field)
	case "alpha":
		return fmt.Sprintf("%s must contain only letters", field)
	case "alphanum":
		return fmt.Sprintf("%s must contain only letters and numbers", field)
	case "url":
		return fmt.Sprintf("%s must be a valid URL", field)
	case "valid_role":
		return fmt.Sprintf("%s must be one of: admin, user, moderator, guest", field)
	case "strong_password":
		return fmt.Sprintf("%s must be at least 8 characters long and contain special characters", field)
	case "oneof":
		return fmt.Sprintf("%s must be one of: %s", field, param)
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", field, param)
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", field, param)
	case "lt":
		return fmt.Sprintf("%s must be less than %s", field, param)
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", field, param)
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}

// Helper function to get validated data from context
func GetValidatedData[T any](c *gin.Context) (T, bool) {
	var zero T
	data, exists := c.Get("validated_data")
	if !exists {
		return zero, false
	}
	
	typed, ok := data.(T)
	if !ok {
		return zero, false
	}
	
	return typed, true
}
