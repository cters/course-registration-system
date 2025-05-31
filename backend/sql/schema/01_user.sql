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

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS go_operator (
    operator_id BIGSERIAL PRIMARY KEY,
    operator_hire_date DATE NOT NULL,
    operator_unit_name VARCHAR(255) NOT NULL DEFAULT '',
    user_id BIGINT NOT NULL UNIQUE,
    CONSTRAINT fk_go_operator_go_user FOREIGN KEY (user_id) REFERENCES go_user(user_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS go_student (
    student_id BIGSERIAL PRIMARY KEY,
    student_credit SMALLINT NOT NULL DEFAULT 0,
    student_enrollment_year SMALLINT NOT NULL,
    student_gpa DECIMAL(8,2) NOT NULL DEFAULT 0.0,
    student_department_id SMALLINT NOT NULL,
    user_id BIGINT NOT NULL UNIQUE,
    CONSTRAINT fk_go_student_go_user FOREIGN KEY (user_id) REFERENCES go_user(user_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS go_instructor (
    instructor_id BIGSERIAL PRIMARY KEY,
    instructor_department_id SMALLINT NOT NULL,
    instructor_hire_date DATE NOT NULL,
    instructor_title VARCHAR(255) NOT NULL DEFAULT '',
    user_id BIGINT NOT NULL UNIQUE,
    CONSTRAINT fk_go_instructor_go_user FOREIGN KEY (user_id) REFERENCES go_user(user_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_go_user_updated_at ON go_user;
-- +goose StatementEnd

-- +goose StatementBegin
DROP FUNCTION IF EXISTS update_updated_at_column;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS go_instructor;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS go_student;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS go_operator;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS go_user;
-- +goose StatementEnd