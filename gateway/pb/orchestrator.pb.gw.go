// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: orchestrator.proto

/*
Package pb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package pb

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_Orchestrator_SendMessage_0(ctx context.Context, marshaler runtime.Marshaler, client OrchestratorClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SendMessageRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.SendMessage(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Orchestrator_SendMessage_0(ctx context.Context, marshaler runtime.Marshaler, server OrchestratorServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SendMessageRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.SendMessage(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterOrchestratorHandlerServer registers the http handlers for service Orchestrator to "mux".
// UnaryRPC     :call OrchestratorServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterOrchestratorHandlerFromEndpoint instead.
func RegisterOrchestratorHandlerServer(ctx context.Context, mux *runtime.ServeMux, server OrchestratorServer) error {

	mux.Handle("POST", pattern_Orchestrator_SendMessage_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/orchestrator.Orchestrator/SendMessage", runtime.WithHTTPPathPattern("/send"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Orchestrator_SendMessage_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Orchestrator_SendMessage_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterOrchestratorHandlerFromEndpoint is same as RegisterOrchestratorHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterOrchestratorHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterOrchestratorHandler(ctx, mux, conn)
}

// RegisterOrchestratorHandler registers the http handlers for service Orchestrator to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterOrchestratorHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterOrchestratorHandlerClient(ctx, mux, NewOrchestratorClient(conn))
}

// RegisterOrchestratorHandlerClient registers the http handlers for service Orchestrator
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "OrchestratorClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "OrchestratorClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "OrchestratorClient" to call the correct interceptors.
func RegisterOrchestratorHandlerClient(ctx context.Context, mux *runtime.ServeMux, client OrchestratorClient) error {

	mux.Handle("POST", pattern_Orchestrator_SendMessage_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/orchestrator.Orchestrator/SendMessage", runtime.WithHTTPPathPattern("/send"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Orchestrator_SendMessage_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Orchestrator_SendMessage_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Orchestrator_SendMessage_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"send"}, ""))
)

var (
	forward_Orchestrator_SendMessage_0 = runtime.ForwardResponseMessage
)
