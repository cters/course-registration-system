package response

const (
	ErrCodeSuccess      = 201 // Success
	ErrCodeParamInvalid = 203 // Email is invalid

	ErrInvalidToken   = 301 // token is invalid
	ErrSendEmailOtp   = 303
	ErrCodeAuthFailed = 405

	ErrCodeInternal      = 500
	ErrCodeUserHasExists = 501 // user has already registered

	ErrCodeSubjectNotFound = 600

	ErrCodeCourseNotFound = 700
)

// message
var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "token is invalid",
	ErrSendEmailOtp:     "Failed to send email OTP",

	ErrCodeUserHasExists: "user has already registered",
	ErrCodeInternal:      "Internal Server Error",
	ErrCodeAuthFailed:    "Authentication failed",
}