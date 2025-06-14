-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS go_course (
        course_id SERIAL PRIMARY KEY,
        subject_id VARCHAR(255) NOT NULL,
        course_term VARCHAR(255) NOT NULL,
        course_max_slot SMALLINT NOT NULL DEFAULT 2000,
        course_current_slot SMALLINT NOT NULL DEFAULT 0,
        course_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        course_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT fk_go_cousre_go_subject FOREIGN KEY (subject_id) REFERENCES go_subject (subject_id) ON DELETE CASCADE
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS go_course;

-- +goose StatementEnd