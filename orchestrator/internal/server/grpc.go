package server

import (
	"context"
	"github.com/tredoc/go-messaging-platform/orchestrator/pb"
)

type OrchestratorGrpcHandler interface {
	SendMessage(context.Context, *pb.SendMessageRequest) (*pb.SendMessageResponse, error)
}

type GRPCServer struct {
	OrchestratorGrpcHandler
}

func NewGRPCServer(h OrchestratorGrpcHandler) *GRPCServer {
	return &GRPCServer{h}
}
