proto:
	protoc --go_out=. --go-grpc_out=. pkg/pb/admin.proto
run:
	go run  cmd/api/main.go 

wire:
	go run  github.com/google/wire/cmd/wire