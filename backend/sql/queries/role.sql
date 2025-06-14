-- name: CreateRole :one
INSERT INTO go_role (role_name, role_created_at, role_updated_at) 
VALUES ($1, NOW(), NOW()) 
RETURNING role_id;

-- name: GetRole :one
SELECT role_id, role_name 
FROM go_role 
WHERE role_id = $1;

-- name: GetRoleByName :one
SELECT role_id, role_name 
FROM go_role 
WHERE role_name = $1;

-- name: UpdateRole :one
UPDATE go_role 
SET role_name = $2, role_updated_at = CURRENT_TIMESTAMP 
WHERE role_id = $1 
RETURNING role_id, role_name;

-- name: UpdateRoleByName :one
UPDATE go_role 
SET role_name = $2, role_updated_at = CURRENT_TIMESTAMP 
WHERE role_name = $1 
RETURNING role_id, role_name;

-- name: RoleExists :one
SELECT EXISTS(SELECT 1 FROM go_role WHERE role_id = $1);