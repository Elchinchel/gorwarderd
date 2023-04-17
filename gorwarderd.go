package gorwarderd

import (
	"fmt"
	"io"
)

const tcpAddress = "localhost:19567"
const unixAddress = "/var/run/forwarderd.sock"

func makeRequest(msgType int, data interface{}) (interface{}, error) {
	rawData, err := dumpFrame(msgType, data)
	if err != nil {
		return nil, err
	}

	conn, err := connect()
	if err != nil {
		return nil, fmt.Errorf("forwarderd: Connect error: %v. Is daemon running?", err)
	}

	_, err = conn.Write(rawData)
	if err != nil {
		return nil, fmt.Errorf("forwarderd: Conn write error: %v", err)
	}

	rawResp, err := io.ReadAll(conn)
	if err != nil {
		return nil, fmt.Errorf("forwarderd: Conn read error: %v", err)
	}

	return loadFrame(rawResp)
}

func SpawnTunnel(data SpawnTunnelRequest) (val SpawnTunnelResponse, err error) {
	resp, err := makeRequest(msgTypeSpawnTunnel, data)
	if err == nil {
		err = loadFrameData(resp, &val)
	}
	return val, err
}
