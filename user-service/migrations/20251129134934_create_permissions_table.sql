-- +goose Up
-- +goose StatementBegin

CREATE TABLE permissions (
     id SERIAL PRIMARY KEY,
     name VARCHAR(150) NOT NULL UNIQUE,
     description TEXT
);

CREATE TABLE role_permissions (
    role_id INTEGER NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    permission_id INTEGER NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS permissions;

-- +goose StatementEnd
