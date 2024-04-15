protoc:
	@echo "Generating protocol buffer files..."
	@protoc --go_out=./golang --go_opt=paths=source_relative \
              --go-grpc_out=./golang --go-grpc_opt=paths=source_relative \
              --grpc-gateway_out=./golang --grpc-gateway_opt paths=source_relative \
              --grpc-gateway_opt generate_unbound_methods=true \
              ./**/*.proto
	@echo "Done."