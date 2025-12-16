package handlers

import (
	"context"
	"log"
	"product-service/databases"
	"time"

	productV1 "github.com/ol1mov-dev/protos/pkg/product/v1"
)

type Product struct {
	Id               int32
	Name             string
	Sku              string
	ShortDescription string
	Description      string
	Price            float32
	PriceOld         float32
	Discount         float32
	Quantity         int32
	IsActive         bool
	Brand            string
	Rating           float32
	CategoryId       int32
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type ProductServer struct {
	productV1.UnimplementedProductV1ServiceServer
}

func (s *ProductServer) CreateProduct(ctx context.Context, req *productV1.CreateProductRequest) (*productV1.CreateProductResponse, error) {
	DB, err := databases.ConnectPostgreSQL()

	if err != nil {
		log.Println("Error connecting to databases")
		return nil, err
	}
	defer DB.Close()

	var productId int32

	query := `INSERT INTO products (name, sku, short_description, description, price, price_old, discount, quantity, is_active, brand, rating, category_id) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`

	log.Println("Creating product...")
	err = DB.QueryRow(
		query,
		req.Name,
		req.Sku,
		req.ShortDescription,
		req.Description,
		req.Price,
		req.PriceOld,
		req.Discount,
		req.Quantity,
		req.IsActive,
		req.Brand,
		req.Rating,
		req.CategoryId,
	).Scan(&productId)

	if err != nil {
		log.Println("Error creating product")
		return nil, err
	}

	log.Println("Product created!")
	return &productV1.CreateProductResponse{ProductId: productId}, nil
}

func (s *ProductServer) UpdateProduct(ctx context.Context, req *productV1.UpdateProductRequest) (*productV1.UpdateProductResponse, error) {
	return &productV1.UpdateProductResponse{}, nil
}

func (s *ProductServer) DeleteProduct(ctx context.Context, req *productV1.DeleteProductRequest) (*productV1.DeleteProductResponse, error) {
	DB, err := databases.ConnectPostgreSQL()

	if err != nil {
		return nil, err
	}

	defer DB.Close()

	var productId int32

	log.Println("Deleting product...")
	query := `DELETE FROM products WHERE id = $1 RETURNING id`
	err = DB.QueryRow(query, req.ProductId).Scan(&productId)

	if err != nil {
		log.Println("Error deleting product")
		return nil, err
	}

	log.Println("Product deleted!")
	return &productV1.DeleteProductResponse{ProductId: productId}, nil
}

func (s *ProductServer) GetAllProductsByQuery(ctx context.Context, req *productV1.GetAllProductsByQueryRequest) (*productV1.GetAllProductsByQueryResponse, error) {
	DB, err := databases.ConnectPostgreSQL()

	if err != nil {
		return nil, err
	}

	defer DB.Close()

	log.Println("Getting all products by query...")
	query := `SELECT id, name, sku, short_description, description, price, price_old, discount, quantity, is_active, brand, rating, category_id
    			FROM products WHERE (is_active = TRUE) AND name LIKE CONCAT('%', $1::text, '%')`

	rows, err := DB.Query(query, req.Query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*productV1.Product

	for rows.Next() {
		product := &productV1.Product{}

		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Sku,
			&product.ShortDescription,
			&product.Description,
			&product.Price,
			&product.PriceOld,
			&product.Discount,
			&product.Quantity,
			&product.IsActive,
			&product.Brand,
			&product.Rating,
			&product.CategoryId,
		)
		if err != nil {
			return nil, err
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	log.Println("Getting products by query done!")
	return &productV1.GetAllProductsByQueryResponse{
		Products: products,
	}, nil
}
