package message

import (
	"errors"
	"time"
)

type Message struct {
	uuid         string
	message      string
	templateUUID string
	sender       string
	receiver     string
	createdAt    time.Time
}

func NewMessage(uuid string, message string, templateUUID string, sender string, receiver string, createdAt time.Time) (Message, error) {
	if uuid == "" {
		return Message{}, errors.New("empty uuid")
	}

	if message == "" {
		return Message{}, errors.New("empty message")
	}

	if templateUUID == "" {
		return Message{}, errors.New("empty template uuid")
	}

	if sender == "" {
		return Message{}, errors.New("empty sender")
	}

	if receiver == "" {
		return Message{}, errors.New("empty receiver")
	}

	return Message{
		uuid:         uuid,
		message:      message,
		templateUUID: templateUUID,
		sender:       sender,
		receiver:     receiver,
		createdAt:    createdAt,
	}, nil
}

func UnmarshalFromDB(uuid string, message string, templateUUID string, sender string, receiver string, createdAt time.Time) (Message, error) {
	return NewMessage(uuid, message, templateUUID, sender, receiver, createdAt)
}

func (m Message) UUID() string {
	return m.uuid
}

func (m Message) Message() string {
	return m.message
}

func (m Message) TemplateUUID() string {
	return m.templateUUID
}

func (m Message) Sender() string {
	return m.sender
}

func (m Message) Receiver() string {
	return m.receiver
}

func (m Message) CreatedAt() time.Time {
	return m.createdAt
}
