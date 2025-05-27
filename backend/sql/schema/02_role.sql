-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS go_role (
        role_id SERIAL PRIMARY KEY,
        role_name VARCHAR,
        role_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        role_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS go_role;

-- +goose StatementEnd