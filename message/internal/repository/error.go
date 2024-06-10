package repository

import "errors"

var (
	ErrMsgNotFound    = errors.New("message not found")
	ErrStatusNotFound = errors.New("status not found")
)
