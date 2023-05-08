package sonyflake_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	. "github.com/tuannguyenandpadcojp/idgenerator/sonyflake"
)

func TestGenerator_NewID(t *testing.T) {
	sf := NewGenerator(time.Date(2023, 5, 8, 0, 0, 0, 0, time.UTC))
	id, err := sf.NewID(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, id)
}

func TestCollision(t *testing.T) {
	const n = 4096

	gen := NewGenerator(time.Date(2023, 5, 8, 0, 0, 0, 0, time.UTC))
	c := make(chan int64, n)
	for i := 0; i < n; i++ {
		go func(c chan int64) {
			id, err := gen.NewID(context.Background())
			require.NoError(t, err)
			c <- id
		}(c)
	}

	// Check for collision.
	m := make(map[int64]struct{})
	for i := 0; i < n; i++ {
		id := <-c
		_, ok := m[id]
		if ok {
			t.Fatalf("collision for %d", id)
		}
		m[id] = struct{}{}
	}
}

// BenchmarkGenerator_NewID
func BenchmarkGenerator_NewID(b *testing.B) {
	sf := NewGenerator(time.Date(2023, 5, 8, 0, 0, 0, 0, time.UTC))
	for i := 0; i < b.N; i++ {
		_, err := sf.NewID(context.Background())
		if err != nil {
			b.Fatal(err)
		}
	}
}
