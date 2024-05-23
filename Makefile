compileProto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pkg/api/auth.proto
build:
	go build cmd/main.go
lint:
	golangci-lint run --fix -v --config golangci.yml
run:
	docker-compose build && docker-compose up
test:
	go test ./...