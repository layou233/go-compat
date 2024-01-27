//go:build !go1.21

package netcompat

import (
	"net"

	"github.com/layou233/go-compat/errorscompat"
)

func ConnMultipathTCP(*net.TCPConn) (bool, error) {
	return false, errorscompat.ErrUnsupported
}

func DialerMultipathTCP(*net.Dialer) bool {
	return false
}

func SetMultipathTCP(*net.Dialer, bool) bool {
	return false
}
