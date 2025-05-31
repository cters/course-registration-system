-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS go_subject (
        subject_id VARCHAR(255) PRIMARY KEY,
        subject_title VARCHAR(255) NOT NULL,
        subject_credit SMALLINT NOT NULL,
        subject_note VARCHAR(255)
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS go_subject;

-- +goose StatementEnd