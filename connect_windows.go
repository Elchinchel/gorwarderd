//go:build windows

package gorwarderd

import "net"

func connect() (net.Conn, error) {
	return net.Dial("tcp", tcpAddress)
}
