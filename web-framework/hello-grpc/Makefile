GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
BRANCH=$(shell git symbolic-ref -q --short HEAD)
REVISION=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shell date +%FT%T%Z)
PROTO_FILES=$(shell find ./helloworld -name *.proto)
PROJECT_PATH=$(shell pwd)

.PHONY: init
init:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

.PHONY: proto
proto:
	protoc  --proto_path=. \
					--proto_path=$(PROJECT_PATH)/helloworld \
					--proto_path=$(PROJECT_PATH)/third_party \
					--go_out=paths=source_relative:. \
					--go-grpc_out=paths=source_relative:. \
					--go-http_out=paths=source_relative:. \
					--go-errors_out=paths=source_relative:. $(PROTO_FILES)

.PHONY: build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION) -X main.Branch=$(BRANCH) -X main.Revision=$(REVISION) -X main.BuildDate=$(BUILD_DATE)" -o ./bin/ ./...

.PHONY: test
test:
	go test -v ./... -cover