
BASE_PROTO_COMMAND = protoc --go_opt=paths=source_relative  --proto_path=../../proto  --go-grpc_opt=paths=source_relative


.PHONY: generate_proto
generate_proto:
	${BASE_PROTO_COMMAND} --go_out=./pkg/grpc  --go-grpc_out=./pkg/grpc ../../proto/inner/v1/processor.proto
