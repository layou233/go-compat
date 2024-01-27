package subtlecompat

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"
)

func TestXORBytes(t *testing.T) {
	for n := 1; n <= 1024; n++ {
		if n > 16 && testing.Short() {
			n += n >> 3
		}
		for alignP := 0; alignP < 8; alignP++ {
			for alignQ := 0; alignQ < 8; alignQ++ {
				for alignD := 0; alignD < 8; alignD++ {
					p := make([]byte, alignP+n, alignP+n+10)[alignP:]
					q := make([]byte, alignQ+n, alignQ+n+10)[alignQ:]
					if n&1 != 0 {
						p = p[:n]
					} else {
						q = q[:n]
					}
					if _, err := io.ReadFull(rand.Reader, p); err != nil {
						t.Fatal(err)
					}
					if _, err := io.ReadFull(rand.Reader, q); err != nil {
						t.Fatal(err)
					}

					d := make([]byte, alignD+n, alignD+n+10)
					for i := range d {
						d[i] = 0xdd
					}
					want := make([]byte, len(d), cap(d))
					copy(want[:cap(want)], d[:cap(d)])
					for i := 0; i < n; i++ {
						want[alignD+i] = p[i] ^ q[i]
					}

					if XORBytes(d[alignD:], p, q); !bytes.Equal(d, want) {
						t.Fatalf("n=%d alignP=%d alignQ=%d alignD=%d:\n\tp = %x\n\tq = %x\n\td = %x\n\twant %x\n", n, alignP, alignQ, alignD, p, q, d, want)
					}
				}
			}
		}
	}
}
