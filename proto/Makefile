USER_PROTO_DIR=user/v1
USER_OUT_DIR=pkg/user/v1

PRODUCT_PROTO_DIR=product/v1
PRODUCT_OUT_DIR=pkg/product/v1

WAREHOUSE_PROTO_DIR=warehouse/v1
WAREHOUSE_OUT_DIR=pkg/warehouse/v1

ORDER_PROTO_DIR=order/v1
ORDER_OUT_DIR=pkg/order/v1

PAYMENT_PROTO_DIR=payment/v1
PAYMENT_OUT_DIR=pkg/payment/v1

.PHONY: proto, clean, user-proto, warehouse-proto, product-proto, order-proto, payment-proto

proto:
	make user-proto
	make product-proto
	make warehouse-proto
	make order-proto
	make payment-proto

user-proto:
	@echo "Generate user folders"
	mkdir -p pkg/user/v1

	@echo "🧩 Generating user gRPC code from proto files..."
	protoc \
		--proto_path=$(USER_PROTO_DIR) \
		--go_out=$(USER_OUT_DIR) \
		--go-grpc_out=$(USER_OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$$(find $(USER_PROTO_DIR) -name "*.proto")
	@echo "✅ Generation complete!"

product-proto:
	@echo "Generate product folders"
	mkdir -p pkg/product/v1

	@echo "🧩 Generating product gRPC code from proto files..."
	protoc \
		--proto_path=$(PRODUCT_PROTO_DIR) \
		--go_out=$(PRODUCT_OUT_DIR) \
		--go-grpc_out=$(PRODUCT_OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$$(find $(PRODUCT_PROTO_DIR) -name "*.proto")
	@echo "✅ Generation complete!"


warehouse-proto:
	@echo "Generate warehouse folders"
	mkdir -p pkg/warehouse/v1

	@echo "🧩 Generating warehouse gRPC code from proto files..."
	protoc \
		--proto_path=$(WAREHOUSE_PROTO_DIR) \
		--go_out=$(WAREHOUSE_OUT_DIR) \
		--go-grpc_out=$(WAREHOUSE_OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$$(find $(WAREHOUSE_PROTO_DIR) -name "*.proto")
	@echo "✅ Generation complete!"

order-proto:
	@echo "Generate order folders"
	mkdir -p pkg/order/v1

	@echo "🧩 Generating order gRPC code from proto files..."
	protoc \
		--proto_path=$(ORDER_PROTO_DIR) \
		--go_out=$(ORDER_OUT_DIR) \
		--go-grpc_out=$(ORDER_OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$$(find $(ORDER_PROTO_DIR) -name "*.proto")
	@echo "✅ Generation complete!"
payment-proto:
	@echo "Generate order folders"
	mkdir -p pkg/payment/v1

	@echo "🧩 Generating order gRPC code from proto files..."
	protoc \
		--proto_path=$(PAYMENT_PROTO_DIR) \
		--go_out=$(PAYMENT_OUT_DIR) \
		--go-grpc_out=$(PAYMENT_OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$$(find $(PAYMENT_PROTO_DIR) -name "*.proto")
	@echo "✅ Generation complete!"

clean:
	rm -rf pkg/user
	rm -rf pkg/product
	rm -rf pkg/order
	rm -rf pkg/warehouse
	rm -rf pkg/payment
