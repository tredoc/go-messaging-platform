package main

import (
	"context"
	"fmt"
	"github.com/tredoc/go-messaging-platform/orchestrator/internal/config"
	"github.com/tredoc/go-messaging-platform/orchestrator/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCServer struct {
	pb.UnimplementedOrchestratorServiceServer
}

func (gs GRPCServer) SendMessage(_ context.Context, _ *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	return &pb.SendMessageResponse{Status: "Message has been sent"}, nil
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
	pb.RegisterOrchestratorServiceServer(grpcServer, &GRPCServer{})

	log.Printf("Starting orchestrator server on port: %s\n", cfg.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
