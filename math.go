package xtypes

import "golang.org/x/exp/constraints"

// Min value
func Min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// Max value
func Max[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}
