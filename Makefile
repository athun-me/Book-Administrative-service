proto:
	protoc --go_out=. --go-grpc_out=. pkg/pb/auth.proto
run:
	go run cmd/api/main.go