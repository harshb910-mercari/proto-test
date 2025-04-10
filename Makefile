PROTO_FILES=$(shell find api -name "*.proto")

generate:
	protoc \
		--proto_path=. \
		--proto_path=$(shell go list -m -f {{.Dir}} github.com/envoyproxy/protoc-gen-validate) \
		--proto_path=$(shell go list -m -f {{.Dir}} github.com/grpc-ecosystem/grpc-gateway/v2) \
		--proto_path=third_party/googleapis \
		--go_out=paths=source_relative:./generated \
		--go-grpc_out=paths=source_relative:./generated \
		--grpc-gateway_out=paths=source_relative:./generated \
		--validate_out="lang=go,paths=source_relative:./generated" \
		$(PROTO_FILES)

clean:
	rm -rf ./generated/api/*.pb.go ./generated/api/*.pb.gw.go ./generated/api/*.validate.go
