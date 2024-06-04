package handler

import (
	"context"
	"github.com/tredoc/go-messaging-platform/message/internal/application"
	"github.com/tredoc/go-messaging-platform/message/internal/application/query"
	"github.com/tredoc/go-messaging-platform/message/pb"
)

type GRPCMessageHandler struct {
	app application.Application
}

func NewGRPCMessageHandler(app application.Application) *GRPCMessageHandler {
	return &GRPCMessageHandler{
		app: app,
	}
}

func (gh GRPCMessageHandler) GetMessageStatus(ctx context.Context, req *pb.GetMessageStatusRequest) (*pb.GetMessageStatusResponse, error) {
	_ = req.GetUuid()

	_ = gh.app.Queries.GetMessageStatus.Handle(ctx, query.GetMessageStatus{})

	return &pb.GetMessageStatusResponse{Status: pb.MessageStatus_NEW}, nil
}
