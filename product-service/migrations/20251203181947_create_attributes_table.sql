-- +goose Up
-- +goose StatementBegin
CREATE TABLE attributes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    data_type VARCHAR(50) NOT NULL DEFAULT 'string' -- например: string, number, bool
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS attributes
-- +goose StatementEnd
