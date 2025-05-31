-- name: GetCourseById :one
SELECT course_id, subject_id, course_term, course_max_slot, course_current_slot FROM go_course WHERE course_id = $1;

-- name: GetCourseBySubjectIdAndTerm :one
SELECT course_id, subject_id, course_term, course_max_slot, course_current_slot FROM go_course WHERE subject_id = $1 AND course_term = $2;

-- name: UpdateCourseById :execresult
UPDATE go_course
SET 
    subject_id = COALESCE($1, subject_id),
    course_term = COALESCE($2, subject_id),
    course_max_slot = COALESCE($3, subject_id),
    course_current_slot = COALESCE($4, subject_id),
    course_updated_at = NOW()
WHERE course_id = $5;

-- name: AddCourse :execresult
INSERT INTO go_course (subject_id, course_term, course_max_slot, course_created_at, course_updated_at) VALUES ($1, $2, $3, NOW(), NOW());

-- name: DeleteCourseById :execresult
DELETE FROM go_course WHERE course_id = $1;

-- name: CheckCourseExist :one
SELECT COUNT(*) FROM go_course WHERE course_id = $1;