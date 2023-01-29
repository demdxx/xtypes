package xtypes

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		prev := 0
		gen := GeneratorSimple(10, func() (int, bool) {
			prev += 1
			return prev, prev <= 10
		})

		list := []int{}
		for v := range gen {
			list = append(list, v)
		}

		assert.Equal(t, 10, len(list))
	})

	t.Run("regular", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		gen := Generator(ctx, 1, func(ctx context.Context) (int, bool) { return 1, true })

		list := []int{<-gen}
		cancel()

		for v := range gen {
			list = append(list, v)
		}

		assert.GreaterOrEqual(t, len(list), 1)
	})
}
