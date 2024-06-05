package server

import (
	"context"
	"github.com/tredoc/go-messaging-platform/template/internal/command"
	"github.com/tredoc/go-messaging-platform/template/internal/query"
	"github.com/tredoc/go-messaging-platform/template/pb"
)

type GRPCHandler struct {
	command command.Command
	query   query.Query
}

func NewGRPCHandler(c command.Command, q query.Query) GRPCHandler {
	return GRPCHandler{
		command: c,
		query:   q,
	}
}

func (gs GRPCHandler) CreateTemplate(_ context.Context, req *pb.CreateTemplateRequest) (*pb.CreateTemplateResponse, error) {
	_ = req.GetType()

	_ = gs.command.CreateTemplate.Handle(context.TODO(), command.CreateTemplate{})

	return &pb.CreateTemplateResponse{Uuid: "New UUID", Status: "created"}, nil
}

func (gs GRPCHandler) GetTemplate(_ context.Context, req *pb.GetTemplateRequest) (*pb.GetTemplateResponse, error) {
	uuid := req.GetUuid()

	_ = gs.query.GetTemplate.Handle(context.TODO(), query.GetTemplate{})

	return &pb.GetTemplateResponse{Uuid: uuid, Template: "Random template string", Type: pb.TemplateType_EMAIL}, nil
}

func (gs GRPCHandler) DeleteTemplate(_ context.Context, req *pb.DeleteTemplateRequest) (*pb.DeleteTemplateResponse, error) {
	uuid := req.GetUuid()

	_ = gs.command.DeleteTemplate.Handle(context.TODO(), command.DeleteTemplate{})

	return &pb.DeleteTemplateResponse{Status: "deleted id: " + uuid}, nil
}
