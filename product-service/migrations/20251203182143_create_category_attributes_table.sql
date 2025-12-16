-- +goose Up
-- +goose StatementBegin
CREATE TABLE category_attributes (
     category_id INT NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
     attribute_id INT NOT NULL REFERENCES attributes(id) ON DELETE CASCADE,
     PRIMARY KEY (category_id, attribute_id)
);

CREATE INDEX idx_category_attributes_category_id ON category_attributes(category_id);
CREATE INDEX idx_category_attributes_attribute_id ON category_attributes(attribute_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS category_attributes;
-- +goose StatementEnd
