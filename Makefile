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
	@go run tools/tools.go

.PHONY: proto/message, proto/orchestrator, proto/template, proto/gateway, proto/all, run/orchestrator, tools