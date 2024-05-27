package main

import (
	"log"
	"os/exec"
)

func main() {
	tools := []string{
		"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest",
		"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
		"google.golang.org/protobuf/cmd/protoc-gen-go@latest",
	}

	for _, tool := range tools {
		cmd := exec.Command("go", "get", tool)
		if err := cmd.Run(); err != nil {
			log.Fatalf("could not get %v: %v", tool, err)
		}
	}
}
