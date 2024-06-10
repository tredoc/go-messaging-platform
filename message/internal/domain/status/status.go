package status

import (
	"errors"
	"time"
)

type MsgStatus int

const (
	NEW       MsgStatus = 0
	SENDING   MsgStatus = 1
	COMPLETED MsgStatus = 2
	FAILED    MsgStatus = 3
)

type Status struct {
	uuid        string
	status      MsgStatus
	messageUUID string
	createdAt   time.Time
}

func NewStatus(uuid string, status MsgStatus, messageUUID string, createdAt time.Time) (Status, error) {
	if uuid == "" {
		return Status{}, errors.New("empty uuid")
	}

	if status < NEW || status > FAILED {
		return Status{}, errors.New("invalid status")
	}

	if messageUUID == "" {
		return Status{}, errors.New("empty message uuid")
	}

	if createdAt.IsZero() {
		return Status{}, errors.New("empty created at")
	}

	return Status{
		uuid:        uuid,
		status:      status,
		messageUUID: messageUUID,
		createdAt:   createdAt,
	}, nil
}

func UnmarshalFromDB(uuid string, status MsgStatus, messageUUID string, createdAt time.Time) (Status, error) {
	return NewStatus(uuid, status, messageUUID, createdAt)
}

func (s Status) UUID() string {
	return s.uuid
}

func (s Status) Status() MsgStatus {
	return s.status
}

func (s Status) MessageUUID() string {
	return s.messageUUID
}

func (s Status) CreatedAt() time.Time {
	return s.createdAt
}
