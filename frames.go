package gorwarderd

import (
	"encoding/json"
	"fmt"
)

type SpawnTunnelRequest struct {
	SshAddr string `json:"ssh_addr"`
	SshUser string `json:"ssh_user"`

	// Must be accessible for daemon
	SshIdentityFile string `json:"ssh_identity_file"`

	RemoteAddr string `json:"remote_addr"`
	LocalAddr  string `json:"local_addr"`
}

type SpawnTunnelResponse struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

const (
	msgTypeSpawnTunnel = 1
	msgTypeRespOk      = 15
	msgTypeRespErr     = 16
)

func dumpFrame(msgType int, data interface{}) ([]byte, error) {
	value, err := json.Marshal([]interface{}{msgType, data})
	if err != nil {
		return value, err
	}
	return append(value, '\n'), err
}

func loadFrameData(data interface{}, v interface{}) error {
	raw, _ := json.Marshal(data)
	return json.Unmarshal(raw, v)
}

func loadFrame(rawData []byte) (interface{}, error) {
	var data []interface{}
	err := json.Unmarshal(rawData[:len(rawData)-1], &data)
	if len(data) != 2 {
		err = fmt.Errorf("forwarderd: Unknown daemon response, bad data type")
	}
	if err != nil {
		return nil, err
	}

	msgType, ok := data[0].(float64)
	if !ok {
		return nil, fmt.Errorf("forwarderd: Unknown daemon response, unknown msg type '%v'", data[0])
	}

	switch int(msgType) {
	case msgTypeRespOk:
		return data[1], nil
	case msgTypeRespErr:
		if respMap, ok := data[1].(map[string]interface{}); !ok {
			return nil, fmt.Errorf("forwarderd: Unknown daemon response, bad data type")
		} else {
			if errString, ok := respMap["error"]; !ok {
				return nil, fmt.Errorf("forwarderd: Unknown daemon response, no error field in error response")
			} else {
				return nil, fmt.Errorf("forwarderd: %s", errString)
			}
		}
	default:
		return nil, fmt.Errorf("forwarderd: Unknown daemon response, unknown msg type '%v'", data[0])
	}
}
