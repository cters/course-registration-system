-- name: GetAllSubject :many
SELECT * FROM go_subject;

-- name: GetSubjectById :one
SELECT * FROM go_subject WHERE subject_id = $1;

-- name: CheckSubjectExist :one
SELECT COUNT(*) FROM go_subject WHERE subject_id = $1;