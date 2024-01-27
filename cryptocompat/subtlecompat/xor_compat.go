//go:build !go1.20

package subtlecompat

import (
	_ "crypto/cipher"
	_ "unsafe"
)

//go:linkname XORBytes crypto/cipher.xorBytes
func XORBytes(dst, x, y []byte) int
