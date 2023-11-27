default: grpc

grpc:
	@protoc -I=protos --go_out=stubs --go-grpc_out=stubs protos/*.proto
