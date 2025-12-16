-- +goose Up
-- +goose StatementBegin

-- admin → все права
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r, permissions p
WHERE r.name = 'admin'
    ON CONFLICT DO NOTHING;

-- seller → товары + просмотр заказов + платежи за свои товары
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
         JOIN permissions p ON p.name IN (
                                          'product.create', 'product.update', 'product.delete', 'product.list',
                                          'order.list',
                                          'payment.list'
    )
WHERE r.name = 'seller'
    ON CONFLICT DO NOTHING;

-- buyer → создание заказов + отмена + просмотр своих заказов
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
         JOIN permissions p ON p.name IN (
                                          'order.create', 'order.cancel', 'order.list',
                                          'payment.create'
    )
WHERE r.name = 'buyer'
    ON CONFLICT DO NOTHING;

-- support → просмотр заказов и пользователей, блокировка пользователей
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
         JOIN permissions p ON p.name IN (
                                          'order.list',
                                          'user.read',
                                          'user.block'
    )
WHERE r.name = 'support'
    ON CONFLICT DO NOTHING;

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DELETE FROM role_permissions;

-- +goose StatementEnd