package xtypes

import "sort"

// Slice type extended with banch of processing methods
type Slice[T any] []T

// SliceApply the function to each element of the slice
func SliceApply[T any, N any](sl []T, apply func(val T) N) Slice[N] {
	nSlice := make(Slice[N], 0, len(sl))
	for _, val := range sl {
		nSlice = append(nSlice, apply(val))
	}
	return nSlice
}

// SliceReduce slice and return new value
func SliceReduce[T any, R any](sl []T, reduce func(val T, ret *R)) R {
	ret := new(R)
	for _, val := range sl {
		reduce(val, ret)
	}
	return *ret
}

// Filter slice values and return new slice without excluded values
func (sl Slice[T]) Filter(filter func(val T) bool) Slice[T] {
	nSlice := make([]T, 0, len(sl))
	for _, val := range sl {
		if filter(val) {
			nSlice = append(nSlice, val)
		}
	}
	return nSlice
}

// Apply the function to each element of the slice
func (sl Slice[T]) Apply(apply func(val T) T) Slice[T] {
	return SliceApply(sl, apply)
}

// ReduceIntoOne slice and return new single value
func (sl Slice[T]) ReduceIntoOne(apply func(val T, ret *T)) T {
	return SliceReduce(sl, apply)
}

// Sort slice values
func (sl Slice[T]) Sort(cmp func(a, b T) bool) Slice[T] {
	sort.Slice(sl, func(i, j int) bool { return cmp(sl[i], sl[j]) })
	return sl
}

// Each iterates every element in the list
func (sl Slice[T]) Each(iter func(val T)) Slice[T] {
	for _, val := range sl {
		iter(val)
	}
	return sl
}
