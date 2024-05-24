proto/message:
	@./scripts/proto.sh message

proto/orchestrator:
	@./scripts/proto.sh orchestrator

proto/template:
	@./scripts/proto.sh template

proto/all:
	@make proto/message
	@make proto/orchestrator
	@make proto/template

.PHONY: proto/message, proto/orchestrator, proto/template, proto/all