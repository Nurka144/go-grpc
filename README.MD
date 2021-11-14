protoc --proto_path=proto proto/\*.proto --go_out=. --go-grpc_out=.

go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc -I . --grpc-gateway_out ./pkg/ --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative proto/adder.proto
