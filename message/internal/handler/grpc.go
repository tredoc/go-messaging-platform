package handler

import (
	"context"
	"fmt"
	"github.com/tredoc/go-messaging-platform/message/pb"
)

type GRPCMessageHandler struct{}

func NewGRPCMessageHandler() *GRPCMessageHandler {
	return &GRPCMessageHandler{}
}

func (gh GRPCMessageHandler) GetMessageStatus(_ context.Context, req *pb.GetMessageStatusRequest) (*pb.GetMessageStatusResponse, error) {
	uuid := req.GetUuid()
	fmt.Println(uuid)

	return &pb.GetMessageStatusResponse{Status: pb.MessageStatus_NEW}, nil
}
