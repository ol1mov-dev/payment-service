-- +goose Up
-- +goose StatementBegin

INSERT INTO permissions (name, description)
VALUES
    -- Товары
    ('product.create', 'Создание товаров'),
    ('product.update', 'Обновление товаров'),
    ('product.delete', 'Удаление товаров'),
    ('product.list',   'Просмотр товаров'),

    -- Заказы
    ('order.create', 'Создание заказов'),
    ('order.update', 'Обновление заказов'),
    ('order.cancel', 'Отмена заказов'),
    ('order.list',   'Просмотр заказов'),

    -- Платежи
    ('payment.create', 'Создание платежей'),
    ('payment.refund', 'Возврат платежей'),
    ('payment.list',   'Просмотр платежей'),

    -- Пользователи
    ('user.read',  'Просмотр пользователей'),
    ('user.block', 'Блокировка пользователей')

    ON CONFLICT (name) DO NOTHING;

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DELETE FROM permissions
WHERE name IN (
               'product.create', 'product.update', 'product.delete', 'product.list',
               'order.create', 'order.update', 'order.cancel', 'order.list',
               'payment.create', 'payment.refund', 'payment.list',
               'user.read', 'user.block'
    );

-- +goose StatementEnd