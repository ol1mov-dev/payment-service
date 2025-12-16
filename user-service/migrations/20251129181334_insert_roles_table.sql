-- +goose Up
-- +goose StatementBegin

INSERT INTO roles (name, description)
VALUES
    ('admin', 'Администратор системы'),
    ('seller', 'Продавец маркетплейса'),
    ('buyer', 'Покупатель'),
    ('support', 'Служба поддержки')
    ON CONFLICT (name) DO NOTHING;

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DELETE FROM roles
WHERE name IN ('admin', 'seller', 'buyer', 'support');

-- +goose StatementEnd