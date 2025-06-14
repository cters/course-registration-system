package model

type CourseInput struct {
	SubjectID         string `json:"subject_id"`
	CourseTerm        string `json:"course_term"`
	CourseMaxSlot     int16  `json:"course_max_slot"`
	CourseCurrentSlot int16  `json:"course_current_slot"`
}

type Course struct {
	CourseID          int32  `json:"course_id"`
	SubjectID         string `json:"subject_id"`
	CourseTerm        string `json:"course_term"`
	CourseMaxSlot     int16  `json:"course_max_slot"`
	CourseCurrentSlot int16  `json:"course_current_slot"`
}
