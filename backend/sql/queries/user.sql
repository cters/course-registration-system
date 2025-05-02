-- name: GetUserByEmail :one
SELECT id, email, username, password, salt FROM go_user WHERE email = $1 LIMIT 1;

-- name: AddUser :execresult
INSERT INTO go_user (
    email, phone, salt, username, password, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, NOW(), NOW()
);

-- name: CheckUserExist :one
SELECT COUNT(*) FROM go_user WHERE email = $1;