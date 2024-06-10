package server

import (
	"context"
	"errors"
	guuid "github.com/google/uuid"
	"github.com/tredoc/go-messaging-platform/message/internal/command"
	"github.com/tredoc/go-messaging-platform/message/internal/domain/message"
	"github.com/tredoc/go-messaging-platform/message/internal/query"
	"github.com/tredoc/go-messaging-platform/message/internal/repository"
	"github.com/tredoc/go-messaging-platform/message/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type GRPCHandler struct {
	messageQuery   query.MessageQuery
	messageCommand command.MessageCommand

	statusQuery   query.StatusQuery
	statusCommand command.StatusCommand
}

func NewGRPCHandler(mq query.MessageQuery, mc command.MessageCommand, sq query.StatusQuery, sc command.StatusCommand) GRPCHandler {
	return GRPCHandler{
		messageQuery:   mq,
		messageCommand: mc,
		statusQuery:    sq,
		statusCommand:  sc,
	}
}

func (gh GRPCHandler) SaveMessage(ctx context.Context, req *pb.SaveMessageRequest) (*pb.SaveMessageResponse, error) {
	msg := req.GetMessage()
	templateUUID := req.GetTemplateUuid()
	sender := req.GetSender()
	receiver := req.GetReceiver()

	dm, err := message.NewMessage(guuid.New().String(), msg, templateUUID, sender, receiver, time.Now())
	if err != nil {
		return nil, err
	}

	err = gh.messageCommand.SaveMessage.Handle(ctx, command.SaveMessage{
		UUID:         dm.UUID(),
		Message:      dm.Message(),
		TemplateUUID: dm.TemplateUUID(),
		Sender:       dm.Sender(),
		Receiver:     dm.Receiver(),
		CreatedAt:    dm.CreatedAt(),
	})

	if err != nil {
		return nil, err
	}

	return &pb.SaveMessageResponse{Uuid: dm.UUID(), CreatedAt: timestamppb.New(dm.CreatedAt())}, nil
}

func (gh GRPCHandler) GetMessageByUUID(ctx context.Context, req *pb.GetMessageByUUIDRequest) (*pb.GetMessageByUUIDResponse, error) {
	uuid := req.GetUuid()

	msg, err := gh.messageQuery.GetMessage.Handle(ctx, query.GetMessage{UUID: uuid})
	if err != nil {
		return nil, err
	}

	resp := &pb.GetMessageByUUIDResponse{
		Uuid:         msg.UUID(),
		Message:      msg.Message(),
		TemplateUuid: msg.TemplateUUID(),
		Sender:       msg.Sender(),
		Receiver:     msg.Receiver(),
	}

	msgStatus, err := gh.statusQuery.GetMessageStatus.Handle(ctx, query.GetMessageStatus{UUID: uuid})
	if err != nil {
		if errors.Is(err, repository.ErrStatusNotFound) {
			resp.CreatedAt = timestamppb.New(msg.CreatedAt())
			resp.Status = pb.MsgStatus_NEW
			resp.UpdatedAt = timestamppb.New(msg.CreatedAt())

			return resp, nil
		}

		return nil, err
	}

	resp.CreatedAt = timestamppb.New(msgStatus.CreatedAt())
	resp.Status = pb.MsgStatus(msgStatus.Status())
	resp.UpdatedAt = timestamppb.New(msgStatus.CreatedAt())

	return resp, nil
}

func (gh GRPCHandler) GetMessageStatusByMessageUUID(ctx context.Context, req *pb.GetMessageStatusByMessageUUIDRequest) (*pb.GetMessageStatusByMessageUUIDResponse, error) {
	messageUUID := req.GetUuid()

	sts, err := gh.statusQuery.GetMessageStatus.Handle(ctx, query.GetMessageStatus{UUID: messageUUID})
	if err != nil {
		return nil, err
	}

	return &pb.GetMessageStatusByMessageUUIDResponse{Uuid: sts.UUID(), MessageUuid: sts.MessageUUID(), Status: pb.MsgStatus(sts.Status()), CreatedAt: timestamppb.New(sts.CreatedAt())}, nil
}
