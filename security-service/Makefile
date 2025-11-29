USER_PROTO_DIR=api/user/v1
USER_OUT_DIR=pkg/api/user/v1

SECURITY_PROTO_DIR=api/security/v1
SECURITY_OUT_DIR=pkg/api/security/v1

PROTOC_GEN_GO=$(shell go env GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC=$(shell go env GOPATH)/bin/protoc-gen-go-grpc

.PHONY: proto clean install-tools


install-tools:
	@echo "⚙️  Installing protoc plugins if missing..."
	@which protoc-gen-go >/dev/null || go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@which protoc-gen-go-grpc >/dev/null || go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

generate-security: install-tools
	@echo "🧩 Generating security gRPC code from proto files..."
	protoc \
		--proto_path=$(SECURITY_PROTO_DIR) \
		--go_out=$(SECURITY_OUT_DIR) \
		--go-grpc_out=$(SECURITY_OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$$(find $(SECURITY_PROTO_DIR) -name "*.proto")
	@echo "✅ Generation complete!"

clean:
	@echo "🧹 Removing generated files..."
	rm -rf $(OUT_DIR)/*.pb.go
