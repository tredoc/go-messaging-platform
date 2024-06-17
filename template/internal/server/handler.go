package server

import (
	"context"
	guuid "github.com/google/uuid"
	"github.com/tredoc/go-messaging-platform/template/internal/command"
	"github.com/tredoc/go-messaging-platform/template/internal/domain/template"
	"github.com/tredoc/go-messaging-platform/template/internal/query"
	"github.com/tredoc/go-messaging-platform/template/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
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

func (gs GRPCHandler) CreateTemplate(ctx context.Context, req *pb.CreateTemplateRequest) (*pb.CreateTemplateResponse, error) {
	content := req.GetContent()
	tmplType := req.GetType()

	t, err := template.NewTemplate(guuid.New().String(), content, template.TemplateType(tmplType), time.Now())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	err = t.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	err = gs.command.CreateTemplate.Handle(ctx, command.CreateTemplate{
		UUID:      t.UUID(),
		Content:   t.Content(),
		TmplType:  t.TmplType(),
		CreatedAt: t.CreatedAt(),
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &pb.CreateTemplateResponse{Uuid: t.UUID()}, nil
}

func (gs GRPCHandler) GetTemplate(ctx context.Context, req *pb.GetTemplateRequest) (*pb.GetTemplateResponse, error) {
	uuid := req.GetUuid()

	t, err := gs.query.GetTemplate.Handle(ctx, query.GetTemplate{UUID: uuid})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &pb.GetTemplateResponse{Uuid: t.UUID(), Content: t.Content(), Type: pb.TemplateType(t.TmplType()), CreatedAt: timestamppb.New(t.CreatedAt())}, nil
}

func (gs GRPCHandler) EnrichTemplate(ctx context.Context, req *pb.EnrichTemplateRequest) (*pb.EnrichTemplateResponse, error) {
	uuid := req.GetUuid()
	message := req.GetMessage()

	if uuid == "" {
		return nil, status.Errorf(codes.InvalidArgument, "uuid of template is required")
	}

	if message == "" {
		return nil, status.Errorf(codes.InvalidArgument, "message is required and cannot be empty")
	}

	tmpl, err := gs.query.GetTemplate.Handle(ctx, query.GetTemplate{UUID: uuid})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	msg := tmpl.Enrich(message)

	return &pb.EnrichTemplateResponse{Message: msg}, nil
}

func (gs GRPCHandler) DeleteTemplate(ctx context.Context, req *pb.DeleteTemplateRequest) (*pb.DeleteTemplateResponse, error) {
	uuid := req.GetUuid()

	err := gs.command.DeleteTemplate.Handle(ctx, command.DeleteTemplate{UUID: uuid})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &pb.DeleteTemplateResponse{Status: "deleted uuid: " + uuid}, nil
}
