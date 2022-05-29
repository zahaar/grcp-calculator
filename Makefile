
generate:
	@mkdir -p gen
	@protoc -I=./pkg/calc --go_out=gen --go_opt=paths=source_relative \
	 --go-grpc_out=gen --go-grpc_opt=paths=source_relative calc.proto
	@echo "==> Generated proto files in ./gen"

dep:
	go get -v -t -d ./...