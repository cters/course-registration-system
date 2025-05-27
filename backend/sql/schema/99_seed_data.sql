-- +goose Up
-- +goose StatementBegin
INSERT INTO
    go_role (role_name)
VALUES
    ('guest'),
    ('student'),
    ('instructor'),
    ('operator');

-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO
    go_permission (permission_name)
VALUES
    ('readOwn'),
    ('readAny'),
    ('createOwn'),
    ('createAny'),
    ('deleteOwn'),
    ('deleteAny'),
    ('updateOwn'),
    ('updateAny');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd