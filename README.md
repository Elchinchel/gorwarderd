## Go client for [forwarderd](https://github.com/elchinchel/forwarderd)


### Usage:
```Go
import (
    "fmt"

    "github.com/elchinchel/gorwarderd"
)

resp, err = SpawnTunnel(SpawnTunnelRequest{
    SshAddr:         "hostname: port",
    SshUser:         "user",
    SshIdentityFile: "/root/.ssh/forward_key", // must be accessible by daemon

    RemoteAddr: "localhost:80",
})

fmt.Printf("%v:%v\n", resp.Host, resp.Port)
```
