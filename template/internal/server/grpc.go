package server

import (
	"context"
	"github.com/tredoc/go-messaging-platform/template/pb"
)

type TemplateGrpcHandler interface {
	CreateTemplate(context.Context, *pb.CreateTemplateRequest) (*pb.CreateTemplateResponse, error)
	GetTemplate(context.Context, *pb.GetTemplateRequest) (*pb.GetTemplateResponse, error)
	DeleteTemplate(context.Context, *pb.DeleteTemplateRequest) (*pb.DeleteTemplateResponse, error)
}

type GRPCServer struct {
	TemplateGrpcHandler
}

func NewGRPCServer(h TemplateGrpcHandler) *GRPCServer {
	return &GRPCServer{h}
}
