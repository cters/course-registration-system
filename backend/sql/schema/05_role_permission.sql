-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS go_role_permission (
        role_id INT,
        permission_id INT,
        role_permission_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        role_permission_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (permission_id, role_id),
        CONSTRAINT fk_go_role_permission_go_role FOREIGN KEY (role_id) REFERENCES go_role (role_id) ON DELETE CASCADE,
        CONSTRAINT fk_go_role_permission_go_permission FOREIGN KEY (permission_id) REFERENCES go_permission (permission_id) ON DELETE CASCADE
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS go_role_permission;

-- +goose StatementEnd