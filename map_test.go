package xtypes

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	var (
		testMap    = map[string]int{"a": 0, "b": 1, "c": 2}
		targetMap1 = map[string]int{"A": 0, "C": 4}
		targetMap2 = map[string]int{"SUM": 5}
	)

	t.Run("Map.Filter.Apply", func(t *testing.T) {
		newMap := Map[string, int](testMap).
			Filter(func(key string, val int) bool { return val%2 == 0 }).
			Apply(func(key string, val int) (string, int) { return strings.ToUpper(key), val * 2 })
		assert.True(t, MapEqual(targetMap1, newMap))
	})

	t.Run("MapReduce", func(t *testing.T) {
		sumMap := MapReduce(testMap, func(key string, val int, ret *map[string]int) {
			if *ret == nil {
				*ret = map[string]int{}
			}
			(*ret)["SUM"] += val * val
		})
		assert.True(t, MapEqual(targetMap2, sumMap))
	})

	t.Run("MapReduce.Sum1", func(t *testing.T) {
		sum := MapReduce(testMap, func(key string, val int, ret *int) { *ret += val })
		assert.Equal(t, 3, sum)
	})

	t.Run("MapReduce.Sum2", func(t *testing.T) {
		sum := Map[string, int](testMap).ReduceIntoOne(func(key string, val int, ret *int) { *ret += val })
		assert.Equal(t, 3, sum)
	})

	t.Run("MapEqual", func(t *testing.T) {
		assert.True(t, MapEqual(map[int]int{1: 1}, map[int]int{1: 1}))
		assert.False(t, MapEqual(map[int]int{1: 1}, map[int]int{}))
		assert.False(t, MapEqual(map[int]int{1: 1}, map[int]int{1: 2}))
	})

	t.Run("Each", func(t *testing.T) {
		assert.Panics(t, func() {
			Map[string, int](testMap).Each(func(k string, a int) { panic("test") })
		})
		assert.Equal(t, len(testMap), len(Map[string, int](testMap).Each(func(k string, a int) {})))
	})

	t.Run("Copy", func(t *testing.T) {
		assert.True(t, MapEqual(testMap, Map[string, int](testMap).Copy()))
	})

	t.Run("Set", func(t *testing.T) {
		assert.True(t, MapEqual(Map[string, int](testMap).Copy().Set("d", 3), map[string]int{"a": 0, "b": 1, "c": 2, "d": 3}))
	})

	t.Run("Keys", func(t *testing.T) {
		assert.ElementsMatch(t, []string{"a", "b", "c"}, Map[string, int](testMap).Keys())
	})

	t.Run("Values", func(t *testing.T) {
		assert.ElementsMatch(t, []int{0, 1, 2}, Map[string, int](testMap).Values())
	})

	t.Run("Merge", func(t *testing.T) {
		assert.True(t, MapEqual(Map[string, int](testMap).Merge(map[string]int{"d": 3}), map[string]int{"a": 0, "b": 1, "c": 2, "d": 3}))
		assert.True(t, MapEqual(Map[string, int](testMap).Merge(map[string]int{"a": 3}), map[string]int{"a": 3, "b": 1, "c": 2}))
		assert.True(t, MapEqual(Map[string, int](testMap).Merge(map[string]int{"a": 3, "b": 4}), map[string]int{"a": 3, "b": 4, "c": 2}))
	})

	t.Run("MergeConflict", func(t *testing.T) {
		assert.True(t, MapEqual(
			Map[string, int](testMap).MergeConflict(
				func(i1, i2 int) int { return i1 + i2 },
				map[string]int{"c": 100, "d": 3}),
			map[string]int{"a": 0, "b": 1, "c": 102, "d": 3}))
	})
}
