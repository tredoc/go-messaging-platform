proto/message:
	@./scripts/proto.sh message

proto/orchestrator:
	@./scripts/proto.sh orchestrator

proto/template:
	@./scripts/proto.sh template

proto/gateway:
	@./scripts/gateway.sh gateway

proto/all:
	@make proto/message
	@make proto/orchestrator
	@make proto/template
	@make proto/gateway

run/orchestrator:
	@go run ./orchestrator/cmd/main.go

tools:
	@echo "Installing gRPC tools"
	@go get google.golang.org/genproto/googleapis/api@v0.0.0-20240528184218-531527333157
	@go get google.golang.org/grpc@v1.64.0
	@go get google.golang.org/protobuf@v1.34.1
	@go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.4
	@echo "Installing gRPC Gateway tools"
	@go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.20.0

dev:
	@docker compose up

.PHONY: proto/message, proto/orchestrator, proto/template, proto/gateway, proto/all, run/orchestrator, tools, dev