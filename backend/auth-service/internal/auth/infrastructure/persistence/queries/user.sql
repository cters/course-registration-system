-- name: GetUserByEmail :one
SELECT user_id, user_password, user_account, user_email, user_name, user_salt, user_phone FROM go_user WHERE user_email = $1 LIMIT 1;

-- name: GetUserById :one
SELECT user_id, user_account, user_email, user_name, user_phone FROM go_user WHERE user_id = $1 LIMIT 1;

-- name: AddUser :one
INSERT INTO go_user (
    user_account, user_email, user_phone, user_salt, user_name, user_password, user_created_at, user_updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, NOW(), NOW()
) RETURNING user_id;

-- name: CheckUserExist :one
SELECT COUNT(*) FROM go_user WHERE user_email = $1;

-- name: CheckStudentExist :one
SELECT COUNT(*) FROM go_student WHERE student_id = $1;

-- name: CheckInstructorExist :one
SELECT COUNT(*) FROM go_instructor WHERE instructor_id = $1;

-- name: CheckOperatorExist :one
SELECT COUNT(*) FROM go_operator WHERE operator_id = $1;