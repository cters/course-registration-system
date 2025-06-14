package model

import "time"

type RegistrationInput struct {
	StudentID  int64  `json:"student_id"`
	CourseID   int32  `json:"course_id"`
	CourseTerm string `json:"course_term"`
}

type RegistrationMessage struct {
	StudentID int64 `json:"student_id"`
	CourseID  int32 `json:"course_id"`
	Timestamp time.Time `json:"timestamp"`
}