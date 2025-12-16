-- +goose Up
-- +goose StatementBegin
CREATE TABLE warehouses
(
    id          BIGSERIAL PRIMARY KEY,
    address     TEXT         NOT NULL,
    city        VARCHAR(100) NOT NULL,
    country     VARCHAR(100) NOT NULL,
    postal_code VARCHAR(20),
    phone       VARCHAR(20),
    is_active   BOOLEAN      NOT NULL DEFAULT TRUE,

    -- Метаданные
    capacity    INTEGER CHECK (capacity >= 0),
    description TEXT,

    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS warehouses;
-- +goose StatementEnd
