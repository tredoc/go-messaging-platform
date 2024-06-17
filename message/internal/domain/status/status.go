package status

import (
	"errors"
	"time"
)

type MessageStatus int

const (
	NEW       MessageStatus = 0
	SENDING   MessageStatus = 1
	COMPLETED MessageStatus = 2
	FAILED    MessageStatus = 3
)

type Status struct {
	uuid        string
	status      MessageStatus
	messageUUID string
	createdAt   time.Time
}

func NewStatus(uuid string, status MessageStatus, messageUUID string, createdAt time.Time) (Status, error) {
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

func UnmarshalFromDB(uuid string, status MessageStatus, messageUUID string, createdAt time.Time) (Status, error) {
	return NewStatus(uuid, status, messageUUID, createdAt)
}

func (s Status) UUID() string {
	return s.uuid
}

func (s Status) Status() MessageStatus {
	return s.status
}

func (s Status) MessageUUID() string {
	return s.messageUUID
}

func (s Status) CreatedAt() time.Time {
	return s.createdAt
}
