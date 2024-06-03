package main

import (
	"context"
	"fmt"
	"github.com/tredoc/go-messaging-platform/template/internal/config"
	"github.com/tredoc/go-messaging-platform/template/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCServer struct {
	pb.UnimplementedTemplateServiceServer
}

func (gs GRPCServer) CreateTemplate(_ context.Context, req *pb.CreateTemplateRequest) (*pb.CreateTemplateResponse, error) {
	template := req.GetTemplate()
	tt := req.GetType()
	fmt.Println(tt, template)

	return &pb.CreateTemplateResponse{Uuid: "New UUID", Status: "created"}, nil
}

func (gs GRPCServer) GetTemplate(_ context.Context, req *pb.GetTemplateRequest) (*pb.GetTemplateResponse, error) {
	uuid := req.GetUuid()
	return &pb.GetTemplateResponse{Uuid: uuid, Template: "Random template string", Type: pb.TemplateType_EMAIL}, nil
}

func (gs GRPCServer) DeleteTemplate(_ context.Context, req *pb.DeleteTemplateRequest) (*pb.DeleteTemplateResponse, error) {
	uuid := req.GetUuid()
	return &pb.DeleteTemplateResponse{Status: "deleted id: " + uuid}, nil
}

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("configuration error: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterTemplateServiceServer(grpcServer, &GRPCServer{})

	log.Printf("Starting template server on port: %s\n", cfg.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
