package response

const (
	ErrCodeSuccess      = 200 // Success
	ErrCodeCreated      = 201 // Created
	ErrCodeAccepted     = 202 // Accepted
	ErrCodeParamInvalid = 203 // Email is invalid

	ErrInvalidToken = 301 // token is invalid
	ErrSendEmailOtp = 303

	ErrCodeBadRequest          = 400
	ErrCodeAuthFailed          = 405
	ErrCodeConflict            = 409
	ErrCodeUnprocessableEntity = 422

	ErrCodeInternal      = 500
	ErrCodeUserHasExists = 501 // user has already registered
	ErrCodeUserNotFound  = 504

	ErrCodeSubjectNotFound = 600

	ErrCodeCourseNotFound = 700
)

// message
var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeCreated:      "Create successful",
	ErrCodeAccepted:     "Request has been accepted and is being processed",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "token is invalid",
	ErrSendEmailOtp:     "Failed to send email OTP",

	ErrCodeBadRequest:          "Bad request",
	ErrCodeAuthFailed:          "Authentication failed",
	ErrCodeConflict:            "Conflict",
	ErrCodeUnprocessableEntity: "Unprocessable Entity",

	ErrCodeInternal:      "Internal Server Error",
	ErrCodeUserHasExists: "user has already registered",
	ErrCodeUserNotFound:  "User not found",

	ErrCodeSubjectNotFound: "Subject not found",
	ErrCodeCourseNotFound:  "Course not found",
}