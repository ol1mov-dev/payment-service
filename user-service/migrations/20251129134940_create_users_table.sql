-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   email VARCHAR(255) NOT NULL UNIQUE,
   password TEXT NOT NULL,

   firstname VARCHAR(100),
   lastname VARCHAR(100),
   phone_number VARCHAR(50),

   role_id INTEGER REFERENCES roles(id) ON DELETE SET NULL,

   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS users;

-- +goose StatementEnd

