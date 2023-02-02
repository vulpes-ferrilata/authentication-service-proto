proto:
	protoc --proto_path=. --go-grpc_out=module=github.com/vulpes-ferrilata/authentication-service-proto:. --go_out=module=github.com/vulpes-ferrilata/authentication-service-proto:. protobuf/*.proto protobuf/models/*.proto

.PHONY: proto