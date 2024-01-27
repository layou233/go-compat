//go:build !go1.21

package compat

import (
	"github.com/layou233/go-compat/cmpcompat"
)

func Max[T cmpcompat.Ordered](x T, y ...T) T {
	maxValue := x
	for _, v := range y {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

func Min[T cmpcompat.Ordered](x T, y ...T) T {
	minValue := x
	for _, v := range y {
		if v < minValue {
			minValue = v
		}
	}
	return minValue
}

func ClearSlice[T ~[]E, E any](t T) {
	var defaultValue E
	for i := range t {
		t[i] = defaultValue
	}
}

type numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func ClearNumericSlice[T ~[]E, E numeric](t T) {
	// https://github.com/golang/go/issues/5373
	for i := range t {
		t[i] = 0
	}
}

func ClearMap[T ~map[K]V, K comparable, V any](t T) {
	for k := range t {
		delete(t, k)
	}
}
