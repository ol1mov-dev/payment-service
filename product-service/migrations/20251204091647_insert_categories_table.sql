-- +goose Up
-- +goose StatementBegin
INSERT INTO categories (name, description) VALUES
   ('Smartphones', 'Mobile phones and accessories'),
   ('Laptops', 'Portable computers and ultrabooks'),
   ('TV', 'Television and home media devices');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
