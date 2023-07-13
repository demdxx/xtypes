package xtypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLazySlice(t *testing.T) {
	sl := NewLazySlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		Filter(func(val int) bool { return val%2 == 0 }).
		Apply(func(val int) int { return val * 2 })

	assert.ElementsMatch(t, Slice[int]{4, 8, 12, 16, 20}, sl.Commit())
}
