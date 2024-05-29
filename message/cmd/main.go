package main

import (
	"context"
	"fmt"
	"github.com/tredoc/go-messaging-platform/message/pb"

	"github.com/tredoc/go-messaging-platform/message/internal/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCServer struct {
	pb.UnimplementedMessageServiceServer
}

func (gs GRPCServer) GetMessageStatus(_ context.Context, req *pb.GetMessageStatusRequest) (*pb.GetMessageStatusResponse, error) {
	id := req.GetId()
	return &pb.GetMessageStatusResponse{Status: fmt.Sprintf("message with id %d is sent", id)}, nil
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
	pb.RegisterMessageServiceServer(grpcServer, &GRPCServer{})

	log.Printf("Starting message server on port: %s\n", cfg.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
