package gorwarderd

import (
	"os"
	"strconv"
	"testing"
)

func getenv(name string, t *testing.T) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		t.Fatalf("Environment variable '%s' not set", name)
	}
	return value
}

func TestSpawn(t *testing.T) {
	localHost := getenv("LOCAL_HOST", t)
	localPort, _ := strconv.Atoi(getenv("LOCAL_PORT", t))

	resp, err := SpawnTunnel(SpawnTunnelRequest{
		SshAddr:         getenv("SSH_ADDR", t),
		SshUser:         getenv("SSH_USER", t),
		SshIdentityFile: getenv("SSH_IDENTITY_FILE", t),

		RemoteAddr: getenv("REMOTE_ADDR", t),
		LocalAddr:  localHost + ":" + strconv.Itoa(localPort),
	})

	if err != nil {
		t.Fatalf("%v", err)
	}

	if resp.Host != localHost || resp.Port != localPort {
		t.Logf("Wanted %v:%v, got %v:%v", localHost, localPort, resp.Host, resp.Port)
		t.Fatalf("Wanted and actual addresses differ")
	}
}

func TestOptionalField(t *testing.T) {
	_, err := SpawnTunnel(SpawnTunnelRequest{
		SshAddr:         getenv("SSH_ADDR", t),
		SshUser:         getenv("SSH_USER", t),
		SshIdentityFile: getenv("SSH_IDENTITY_FILE", t),

		RemoteAddr: getenv("REMOTE_ADDR", t),
		LocalAddr:  "",
	})

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestUnknownMessageType(t *testing.T) {
	_, err := makeRequest(1024, nil)
	if err == nil {
		t.Fatalf("%v", err)
	}
}
