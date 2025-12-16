-- +goose Up
-- +goose StatementBegin
INSERT INTO category_attributes (category_id, attribute_id) VALUES
    ((SELECT id FROM categories WHERE name='Smartphones'),
     (SELECT id FROM attributes WHERE name='Color')),

    ((SELECT id FROM categories WHERE name='Smartphones'),
     (SELECT id FROM attributes WHERE name='BatteryCapacity')),

    ((SELECT id FROM categories WHERE name='Smartphones'),
     (SELECT id FROM attributes WHERE name='ScreenSize')),

    ((SELECT id FROM categories WHERE name='Laptops'),
     (SELECT id FROM attributes WHERE name='Weight')),

    ((SELECT id FROM categories WHERE name='Laptops'),
     (SELECT id FROM attributes WHERE name='ScreenSize')),

    ((SELECT id FROM categories WHERE name='TV'),
     (SELECT id FROM attributes WHERE name='ScreenSize'));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
