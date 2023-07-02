gen:
	protoc --go_out=gen/go/static/v1 --go-grpc_out=gen/go/static/v1 --proto_path=proto/static/v1 --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative service.proto