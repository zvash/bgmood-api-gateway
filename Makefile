test:
	go test -v -cover ./...

server:
	go run main.go

proto-api:
	rm -rf internal/pb/*.go 2>/dev/null
	protoc --go_out=internal/pb --go_opt=paths=source_relative --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative --grpc-gateway_out=internal/pb --grpc-gateway_opt=paths=source_relative --proto_path=internal/proto internal/proto/*.proto


proto-auth:
	rm -rf internal/authpb/* 2>/dev/null
	protoc --go_out=internal/authpb --go_opt=paths=source_relative --go-grpc_out=internal/authpb --go-grpc_opt=paths=source_relative --grpc-gateway_out=internal/authpb --grpc-gateway_opt=paths=source_relative --proto_path=../bgmood-protos/auth-service ../bgmood-protos/auth-service/*.proto

proto-file:
	rm -rf internal/filepb/* 2>/dev/null
	protoc --go_out=internal/filepb --go_opt=paths=source_relative --go-grpc_out=internal/filepb --go-grpc_opt=paths=source_relative --grpc-gateway_out=internal/filepb --grpc-gateway_opt=paths=source_relative --proto_path=../bgmood-protos/file-service ../bgmood-protos/file-service/*.proto

proto:
	make proto-api
	make proto-auth
	make proto-file

.PHONY: test server proto proto-auth proto-api proto-file