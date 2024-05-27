package main

import (
	"fmt"
	pb "github.com/tredoc/go-grpc/proto/gen"
	"github.com/tredoc/go-messaging-platform/orchestrator/internal/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Responser struct {
	pb.UnimplementedResponserServer
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
	pb.RegisterResponserServer(grpcServer, &Responser{})

	log.Printf("Starting orchestrator server on port: %s\n", cfg.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
