-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS go_subject_requirement (
        subject_id VARCHAR(255) NOT NULL,
        requirement_id VARCHAR(255) NOT NULL,
        requirement_type VARCHAR(255) CHECK (
            requirement_type IN ('prerequisite', 'corequisite', 'recommended')
        ),
        PRIMARY KEY (subject_id, requirement_id),
        CONSTRAINT fk_go_subject_requirement_subject_id FOREIGN KEY (subject_id) REFERENCES go_subject (subject_id) ON DELETE CASCADE,
        CONSTRAINT fk_go_subject_requirement_requirement_id FOREIGN KEY (requirement_id) REFERENCES go_subject (subject_id) ON DELETE CASCADE
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS go_subject_requirement;

-- +goose StatementEnd