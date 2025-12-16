-- +goose Up
-- +goose StatementBegin
INSERT INTO attributes (name, data_type) VALUES
     ('Color', 'string'),
     ('Weight', 'number'),
     ('IsNew', 'bool'),
     ('ScreenSize', 'number'),
     ('BatteryCapacity', 'number');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
