package xtypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	var (
		testSlice        = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		testSliceOrdered = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
		targetSlice1     = []int{0, 4, 8, 12, 16, 20}
		targetSlice2     = []int{0, 16, 64, 144, 256, 400}
	)

	t.Run("Slice.Filter.Apply", func(t *testing.T) {
		newSlcie := Slice[int](testSlice).
			Filter(func(val int) bool { return val%2 == 0 }).
			Apply(func(val int) int { return val * 2 })
		assert.ElementsMatch(t, targetSlice1, newSlcie)
	})

	t.Run("SliceReduce::square", func(t *testing.T) {
		squareSlice := SliceReduce(targetSlice1, func(val int, ret *[]int) { *ret = append(*ret, val*val) })
		assert.ElementsMatch(t, targetSlice2, squareSlice)
	})

	t.Run("SliceReduce::sum1", func(t *testing.T) {
		sum := SliceReduce(testSlice, func(val int, ret *int) { *ret += val })
		assert.Equal(t, 1+2+3+4+5+6+7+8+9+10, sum)
	})

	t.Run("SliceReduce::sum2", func(t *testing.T) {
		sum := Slice[int](testSlice).ReduceIntoOne(func(val int, ret *int) { *ret += val })
		assert.Equal(t, 1+2+3+4+5+6+7+8+9+10, sum)
	})

	t.Run("Sort", func(t *testing.T) {
		ordered := Slice[int](append([]int{}, testSlice...)).Sort(func(a, b int) bool { return a > b })
		assert.ElementsMatch(t, testSliceOrdered, ordered)
	})

	t.Run("Each", func(t *testing.T) {
		assert.Panics(t, func() {
			Slice[int](testSlice).Each(func(a int) { panic("test") })
		})
		assert.Equal(t, len(testSlice), len(Slice[int](testSlice).Each(func(a int) {})))
	})

	t.Run("IndexOf", func(t *testing.T) {
		assert.Equal(t, 5, Slice[int](testSlice).IndexOf(func(a int) bool { return a == 5 }))
		assert.Equal(t, -1, Slice[int](testSlice).IndexOf(func(a int) bool { return a == 11 }))
	})

	t.Run("BinarySearch", func(t *testing.T) {
		assert.Equal(t, 5, Slice[int](testSlice).BinarySearch(func(a int) bool { return a == 5 }))
		assert.Equal(t, -1, Slice[int](testSlice).BinarySearch(func(a int) bool { return a == 11 }))
	})

	t.Run("Append", func(t *testing.T) {
		assert.ElementsMatch(t, append(testSlice, 11), Slice[int](testSlice).Append(11))
	})

	t.Run("Prepend", func(t *testing.T) {
		assert.ElementsMatch(t, append([]int{11}, testSlice...), Slice[int](testSlice).Prepend(11))
	})

	t.Run("ValueOr", func(t *testing.T) {
		assert.Equal(t, 1, Slice[int](testSlice).ValueOr(1, 0))
		assert.Equal(t, 1, Slice[int](testSlice).ValueOr(11, 1))
	})

	t.Run("RemoveAt", func(t *testing.T) {
		sl := Slice[int](append([]int{}, testSlice...))
		s1 := sl.RemoveAt(5)
		assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 6, 7, 8, 9, 10}, s1)
		s2 := s1.RemoveAt(11)
		assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 6, 7, 8, 9, 10}, s2)
	})

	t.Run("RemoveRange", func(t *testing.T) {
		sl := Slice[int](append([]int{}, testSlice...))
		s1 := sl.RemoveRange(10, 8)
		assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 10}, s1)
		s2 := s1.RemoveRange(8, 12)
		assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 5, 6, 7}, s2)
		s3 := s2.RemoveRange(100, 200)
		assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 5, 6, 7}, s3)
	})

	t.Run("Copy", func(t *testing.T) {
		assert.ElementsMatch(t, testSlice, Slice[int](testSlice).Copy())
	})
}
