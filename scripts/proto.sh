#!/bin/bash
set -e

readonly service="$1"

rm -rf "$service/pb"
mkdir -p "$service/pb"

protoc \
  --proto_path=api/protobuf "api/protobuf/$service.proto" \
  "--go_out=$service/pb" --go_opt=paths=source_relative \
  --go-grpc_opt=require_unimplemented_servers=false \
  "--go-grpc_out=$service/pb" --go-grpc_opt=paths=source_relative

cp -r $service/pb/* gateway/pb