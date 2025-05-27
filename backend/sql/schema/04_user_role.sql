-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS go_user_role (
        user_id INT,
        role_id INT,
        user_role_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        user_role_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (user_id, role_id),
        CONSTRAINT fk_go_user_role_go_user FOREIGN KEY (user_id) REFERENCES go_user (user_id) ON DELETE CASCADE,
        CONSTRAINT fk_go_user_role_go_role FOREIGN KEY (role_id) REFERENCES go_role (role_id) ON DELETE CASCADE
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS go_user_role;

-- +goose StatementEnd