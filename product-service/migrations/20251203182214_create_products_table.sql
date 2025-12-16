-- +goose Up
-- +goose StatementBegin
CREATE TABLE products (
      id SERIAL PRIMARY KEY,
      name VARCHAR(255) NOT NULL,
      sku VARCHAR(100) NOT NULL UNIQUE,
      short_description VARCHAR(500),
      description TEXT,
      price NUMERIC(10, 2) NOT NULL DEFAULT 0,
      price_old NUMERIC(10, 2),
      discount NUMERIC(5, 2),
      quantity INT NOT NULL DEFAULT 0,
      is_active BOOLEAN NOT NULL DEFAULT TRUE,
      brand VARCHAR(255),
      rating NUMERIC(3, 2) DEFAULT 0, -- например 4.7
      category_id INT NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
      created_at TIMESTAMP NOT NULL DEFAULT NOW(),
      updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_products_category_id ON products(category_id);
CREATE INDEX idx_products_is_active ON products(is_active);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
