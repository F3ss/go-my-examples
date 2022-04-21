.PHONY: run
run:
	go run cmd/app/main.go

.PHONY: proto
proto:
	protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/adder.proto