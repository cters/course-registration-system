-- name: GetUserByEmail :one
SELECT user_id, user_email, user_name, user_password, user_salt FROM go_user WHERE user_email = $1 LIMIT 1;

-- name: AddUser :execresult
INSERT INTO go_user (
    user_account, user_email, user_phone, user_salt, user_name, user_password, user_created_at, user_updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, NOW(), NOW()
);

-- name: CheckUserExist :one
SELECT COUNT(*) FROM go_user WHERE user_email = $1;