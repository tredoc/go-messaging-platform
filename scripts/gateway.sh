#!/bin/bash
set -e

readonly service="$1"

protoc -I ./api/protobuf "--grpc-gateway_out=$service/pb" \
    --grpc-gateway_opt paths=source_relative \
    api/protobuf/*.proto