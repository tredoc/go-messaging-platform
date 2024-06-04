package handler

import (
	"context"
	"github.com/tredoc/go-messaging-platform/orchestrator/pb"
)

type GRPCOrchestratorHandler struct{}

func NewGRPCOrchestratorHandler() *GRPCOrchestratorHandler {
	return &GRPCOrchestratorHandler{}
}

func (gs GRPCOrchestratorHandler) SendMessage(_ context.Context, _ *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	return &pb.SendMessageResponse{Status: pb.OrchestratorMessageStatus_CREATED}, nil
}
