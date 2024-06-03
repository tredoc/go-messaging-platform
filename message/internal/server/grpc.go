package server

import (
	"context"
	"github.com/tredoc/go-messaging-platform/message/pb"
)

type MessageGrpcHandler interface {
	GetMessageStatus(context.Context, *pb.GetMessageStatusRequest) (*pb.GetMessageStatusResponse, error)
}

type GRPCServer struct {
	MessageGrpcHandler
}

func NewGRPCServer(h MessageGrpcHandler) *GRPCServer {
	return &GRPCServer{h}
}
