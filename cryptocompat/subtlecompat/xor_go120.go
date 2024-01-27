//go:build go1.20

package subtlecompat

import (
	_ "crypto/subtle"
	_ "unsafe"
)

// XORBytes sets dst[i] = x[i] ^ y[i] for all i < n = min(len(x), len(y)),
// returning n, the number of bytes written to dst.
// If dst does not have length at least n,
// XORBytes panics without writing anything to dst.
//
//go:linkname XORBytes crypto/subtle.XORBytes
func XORBytes(dst, x, y []byte) int
