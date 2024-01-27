//go:build go1.21

package compat

import "cmp"

func Max[T cmp.Ordered](x T, y ...T) T {
	return max(x, y...)
}

func Min[T cmp.Ordered](x T, y ...T) T {
	return min(x, y...)
}

func ClearSlice[T ~[]E, E any](t T) {
	clear(t)
}

type numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// ClearNumericSlice is like ClearSlice, but only receives slices of
// numeric types. This is expected to have better performance
// when compiled with legacy Go versions, with optimized compatible code.
func ClearNumericSlice[T ~[]E, E numeric](t T) {
	clear(t)
}

func ClearMap[T ~map[K]V, K comparable, V any](t T) {
	clear(t)
}
