.PHONY: build
build:
	go build -ldflags="-s -w" -o bin/app cmd/app/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: cover
cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: protogen
protogen:
	protoc --go_out=.  --go-grpc_out=. pkg/proto/*.proto
