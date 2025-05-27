-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS go_user (
    user_id SERIAL PRIMARY KEY, 
    user_account VARCHAR(255) NOT NULL DEFAULT '' UNIQUE,
    user_password VARCHAR(255) NOT NULL DEFAULT '', 
    user_email VARCHAR(255) NOT NULL DEFAULT '' UNIQUE, 
    user_name VARCHAR(255) NOT NULL DEFAULT '', 
    user_phone VARCHAR(255) NOT NULL DEFAULT '' UNIQUE, 
    user_credit SMALLINT NOT NULL DEFAULT 0,
    user_salt VARCHAR(255) NOT NULL DEFAULT '',
    user_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    user_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER update_go_user_updated_at
BEFORE UPDATE ON go_user
FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_go_user_updated_at ON go_user;
-- +goose StatementEnd

-- +goose StatementBegin
DROP FUNCTION IF EXISTS update_updated_at_column;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS go_user;
-- +goose StatementEnd