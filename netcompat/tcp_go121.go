//go:build go1.21

package netcompat

import "net"

func ConnMultipathTCP(conn *net.TCPConn) (bool, error) {
	return conn.MultipathTCP()
}

func DialerMultipathTCP(dialer *net.Dialer) bool {
	return dialer.MultipathTCP()
}

func SetMultipathTCP(dialer *net.Dialer, use bool) bool {
	dialer.SetMultipathTCP(use)
	return true
}
