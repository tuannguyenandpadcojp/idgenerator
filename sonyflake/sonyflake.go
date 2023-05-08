package sonyflake

import (
	"context"
	"time"

	"github.com/sony/sonyflake"
)

type Generator struct {
	sf *sonyflake.Sonyflake
}

func NewGenerator(startTime time.Time) *Generator {
	return &Generator{
		sf: sonyflake.NewSonyflake(sonyflake.Settings{StartTime: startTime}),
	}
}

// NewID implements idgen.IDGenerator using https://github.com/sony/sonyflake
// The params are ignored.
func (g *Generator) NewID(_ context.Context) (int64, error) {
	id, err := g.sf.NextID()
	if err != nil {
		return 0, err
	}

	// Because A Sonyflake ID is 63 bits, so we can safely convert it into int64.
	return int64(id), nil
}
