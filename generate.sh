#protoc --go_out=gen --go_opt=paths=source_relative --go-grpc_out=gen --go-grpc_opt=paths=source_relative proto/**/*.proto
protoc --go_out=gen \
	   --go_opt=module=github.com/fleanza74/test-proto/proto \
	   --go-grpc_out=gen \
	   --go-grpc_opt=module=github.com/fleanza74/test-proto/proto \
	   proto/**/*.proto

