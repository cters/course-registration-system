-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS go_registration (
        registration_id BIGSERIAL PRIMARY KEY,
        student_id BIGINT NOT NULL,
        course_id INT NOT NULL,
        registration_status VARCHAR(255) NOT NULL,
        registration_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT fk_go_registration_go_student FOREIGN KEY (student_id) REFERENCES go_student (student_id) ON DELETE CASCADE,
        CONSTRAINT fk_go_registration_go_course FOREIGN KEY (course_id) REFERENCES go_course (course_id) ON DELETE CASCADE
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS go_registration;

-- +goose StatementEnd