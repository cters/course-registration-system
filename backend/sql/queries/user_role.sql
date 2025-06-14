-- name: AssignRoleToUser :execresult
INSERT INTO go_user_role (user_id, role_id) 
VALUES ($1, $2);

-- name: RemoveRoleFromUser :exec
DELETE FROM go_user_role 
WHERE user_id = $1 AND role_id = $2;

-- name: GetUserRoles :many
SELECT ur.user_id, ur.role_id, r.role_name
FROM go_user_role ur
JOIN go_role r ON ur.role_id = r.role_id
WHERE ur.user_id = $1
ORDER BY r.role_name;

-- name: GetRoleUsers :many
SELECT ur.user_id, ur.role_id, u.user_name, u.user_email
FROM go_user_role ur
JOIN go_user u ON ur.user_id = u.user_id
WHERE ur.role_id = $1
ORDER BY u.user_name;

-- name: GetUserRolesByUserName :many
SELECT ur.user_id, ur.role_id, r.role_name
FROM go_user_role ur
JOIN go_role r ON ur.role_id = r.role_id
JOIN go_user u ON ur.user_id = u.user_id
WHERE u.user_name = $1
ORDER BY r.role_name;