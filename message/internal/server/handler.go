package server

import (
	"context"
	"github.com/tredoc/go-messaging-platform/message/internal/query"
	"github.com/tredoc/go-messaging-platform/message/pb"
)

type GRPCHandler struct {
	query query.Query
}

func NewGRPCHandler(q query.Query) GRPCHandler {
	return GRPCHandler{
		query: q,
	}
}

func (gh GRPCHandler) GetMessageStatus(ctx context.Context, req *pb.GetMessageStatusRequest) (*pb.GetMessageStatusResponse, error) {
	_ = req.GetUuid()

	_ = gh.query.GetMessageStatus.Handle(ctx, query.GetMessageStatus{})

	return &pb.GetMessageStatusResponse{Status: pb.MsgStatus_NEW}, nil
}
