-- name: CreateRegistration :execresult
INSERT INTO go_registration (student_id, course_id, registration_status) VALUES ($1, $2, $3);

-- name: GetStudentRegistrations :many
SELECT r.*, c.subject_id, c.course_term
FROM go_registration r
JOIN go_course c ON r.course_id = c.course_id
WHERE r.student_id = $1
ORDER BY r.registration_date DESC;

-- name: GetCourseRegistrations :many
SELECT r.*, s.student_id, u.user_name, u.user_email
FROM go_registration r
JOIN go_student s ON r.student_id = s.student_id
JOIN go_user u ON s.user_id = u.user_id
WHERE r.course_id = $1
ORDER BY r.registration_date;

-- name: UpdateRegistrationStatus :one
UPDATE go_registration 
SET registration_status = $2
WHERE registration_id = $1
RETURNING *;

-- name: DeleteRegistration :execresult
DELETE FROM go_registration WHERE registration_id = $1;