package server

import (
	"context"
	"github.com/tredoc/go-messaging-platform/orchestrator/pb"
)

type GRPCHandler struct {
	msgClient  pb.MessageServiceClient
	tmplClient pb.TemplateServiceClient
}

func NewGRPCHandler(msgClient pb.MessageServiceClient, tmplClient pb.TemplateServiceClient) GRPCHandler {

	return GRPCHandler{
		msgClient:  msgClient,
		tmplClient: tmplClient,
	}
}

func (gs GRPCHandler) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	msg := req.GetMessage()
	templateUUID := req.GetTemplateUuid()
	sender := req.GetSender()
	receiver := req.GetReceiver()

	tmplResp, err := gs.tmplClient.EnrichTemplate(ctx, &pb.EnrichTemplateRequest{
		Uuid:    templateUUID,
		Message: msg,
	})
	if err != nil {
		return nil, err
	}

	msgResp, err := gs.msgClient.SaveMessage(ctx, &pb.SaveMessageRequest{
		Message:      tmplResp.GetMessage(),
		Sender:       sender,
		Receiver:     receiver,
		TemplateUuid: templateUUID,
	})
	if err != nil {
		return nil, err
	}

	return &pb.SendMessageResponse{Uuid: msgResp.GetUuid(), Status: pb.OrchestratorMessageStatus_NEW, CreatedAt: msgResp.GetCreatedAt()}, nil
}
