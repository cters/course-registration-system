-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS go_permission (
        permission_id SERIAL PRIMARY KEY,
        permission_name VARCHAR,
        permission_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        permission_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS go_permission;

-- +goose StatementEnd