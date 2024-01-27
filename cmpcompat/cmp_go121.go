//go:build go1.21

package cmpcompat

import (
	"cmp"
)

type Ordered = cmp.Ordered

// Less reports whether x is less than y.
// For floating-point types, a NaN is considered less than any non-NaN,
// and -0.0 is not less than (is equal to) 0.0.
// See cmp.Less for more details.
func Less[T Ordered](x, y T) bool {
	return cmp.Less(x, y)
}

// Compare returns
//
//	-1 if x is less than y,
//	 0 if x equals y,
//	+1 if x is greater than y.
//
// For floating-point types, a NaN is considered less than any non-NaN,
// a NaN is considered equal to a NaN, and -0.0 is equal to 0.0.
// See cmp.Compare for more details.
func Compare[T Ordered](x, y T) int {
	return cmp.Compare(x, y)
}
