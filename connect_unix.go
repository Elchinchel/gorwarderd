//go:build darwin || linux

package gorwarderd

import "net"

func connect() (net.Conn, error) {
	conn, err := net.Dial("unix", unixAddress)
	if err != nil {
		conn, err = net.Dial("tcp", tcpAddress)
	}
	return conn, err
}
