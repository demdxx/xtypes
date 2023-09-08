package xtypes

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		prev := 0
		gen := GeneratorSimple(10, func(_ int) (int, bool) {
			prev += 1
			return prev, prev <= 10
		})

		list := []int{}
		for v := range gen {
			list = append(list, v)
		}

		assert.Equal(t, 10, len(list))
	})

	regularTestFn := func(stop1, stop2 int, cancelf context.CancelFunc) func(_ context.Context, _ int) (int, bool) {
		return func(ctx context.Context, prev int) (int, bool) {
			switch prev {
			case stop1:
				return 0, false
			case stop2:
				cancelf()
				return 0, true
			}
			return prev + 1, true
		}
	}

	regularTest := func(ctx context.Context, size int, fn func(_ context.Context, _ int) (int, bool)) func(t *testing.T) {
		return func(t *testing.T) {
			gen := Generator(ctx, size, fn)

			list := []int{}
			for v := range gen {
				list = append(list, v)
			}

			for v := range gen {
				list = append(list, v)
			}

			assert.GreaterOrEqual(t, len(list), 1)
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Run("regular1", regularTest(ctx, 1, regularTestFn(1, 2, cancel)))
	cancel()

	ctx, cancel = context.WithCancel(context.Background())
	t.Run("regular2", regularTest(ctx, 2, regularTestFn(1, 2, cancel)))
	cancel()

	ctx, cancel = context.WithCancel(context.Background())
	t.Run("regular3", regularTest(ctx, 3, regularTestFn(-1, 2, cancel)))
	cancel()
}
