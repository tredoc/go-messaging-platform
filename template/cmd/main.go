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

func (gs GRPCServer) GetTemplateByID(_ context.Context, req *pb.GetTemplateByIDRequest) (*pb.GetTemplateByIDResponse, error) {
	id := req.GetId()
	return &pb.GetTemplateByIDResponse{Name: "Random Template ID: " + id}, nil
}

func (gs GRPCServer) CreateTemplate(_ context.Context, in *pb.CreateTemplateRequest) (*pb.CreateTemplateResponse, error) {
	fmt.Println(in)
	return &pb.CreateTemplateResponse{Id: "new uuid"}, nil
}

func (gs GRPCServer) DeleteTemplateByID(_ context.Context, req *pb.DeleteTemplateByIDRequest) (*pb.DeleteTemplateByIDResponse, error) {
	id := req.GetId()
	return &pb.DeleteTemplateByIDResponse{Id: "deleted id: " + id}, nil
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
