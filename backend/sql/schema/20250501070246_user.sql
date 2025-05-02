-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS go_user (
    id SERIAL PRIMARY KEY, 
    email VARCHAR(255) NOT NULL DEFAULT '' UNIQUE, 
    phone VARCHAR(255) NOT NULL DEFAULT '' UNIQUE, 
    username VARCHAR(255) NOT NULL DEFAULT '', 
    password VARCHAR(255) NOT NULL DEFAULT '', 
    salt VARCHAR(255) NOT NULL DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
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