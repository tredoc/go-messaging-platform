package status

import (
	"context"
)

type Repository interface {
	SaveStatus(context.Context, Status) error
	FindStatusByUUID(context.Context, string) (Status, error)
}
