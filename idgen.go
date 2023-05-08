package idgenerator

import "context"

type IDGenerator interface {
	NewID(ctx context.Context) (int64, error)
}
