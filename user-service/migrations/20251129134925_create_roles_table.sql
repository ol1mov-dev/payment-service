-- +goose Up
-- +goose StatementBegin

CREATE TABLE roles (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(100) NOT NULL UNIQUE,
                       description TEXT
);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS roles;

-- +goose StatementEnd

