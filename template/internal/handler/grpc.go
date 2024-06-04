package handler

import (
	"context"
	application "github.com/tredoc/go-messaging-platform/template/internal/application"
	"github.com/tredoc/go-messaging-platform/template/internal/application/command"
	"github.com/tredoc/go-messaging-platform/template/internal/application/query"
	"github.com/tredoc/go-messaging-platform/template/pb"
)

type GRPCTemplateHandler struct {
	app application.Application
}

func NewGRPCTemplateHandler(app application.Application) GRPCTemplateHandler {
	return GRPCTemplateHandler{
		app: app,
	}
}

func (gs GRPCTemplateHandler) CreateTemplate(_ context.Context, req *pb.CreateTemplateRequest) (*pb.CreateTemplateResponse, error) {
	_ = req.GetType()

	_ = gs.app.Commands.CreateTemplate.Handle(context.TODO(), command.CreateTemplate{})

	return &pb.CreateTemplateResponse{Uuid: "New UUID", Status: "created"}, nil
}

func (gs GRPCTemplateHandler) GetTemplate(_ context.Context, req *pb.GetTemplateRequest) (*pb.GetTemplateResponse, error) {
	uuid := req.GetUuid()

	_ = gs.app.Queries.GetTemplate.Handle(context.TODO(), query.GetTemplate{})

	return &pb.GetTemplateResponse{Uuid: uuid, Template: "Random template string", Type: pb.TemplateType_EMAIL}, nil
}

func (gs GRPCTemplateHandler) DeleteTemplate(_ context.Context, req *pb.DeleteTemplateRequest) (*pb.DeleteTemplateResponse, error) {
	uuid := req.GetUuid()

	_ = gs.app.Commands.DeleteTemplate.Handle(context.TODO(), command.DeleteTemplate{})

	return &pb.DeleteTemplateResponse{Status: "deleted id: " + uuid}, nil
}
