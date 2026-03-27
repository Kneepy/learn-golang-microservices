.PHONY: proto run

proto:
	protoc --proto_path=proto --go_out=proto/pb --go_opt=paths=source_relative --go-grpc_out=proto/pb --go-grpc_opt=paths=source_relative proto/*.proto

# для теста
	protoc --proto_path=proto --go_out=gen --go_opt=paths=source_relative --go-grpc_out=gen --go-grpc_opt=paths=source_relative proto/user/*.proto
run:
	go run ./cmd/server
