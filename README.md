# Middle Man Service (M2S)



## Get Started

### Install dependency

```shell
go get github.com/opposite-bracket/m2s
```

### Run m2s to log all requests coming in

```go
// main.go
package main

import (
	"github.com/opposite-bracket/m2s"
	"log"
)

func main() {
	srv, err := m2s.NewService()
	if err != nil {
		log.Panicf(
			"failed to instantiate m2 service: [error: %s]",
			err,
		)
	}

	if err := srv.RunServer(); err != nil {
		log.Fatalf("server errored: %v", err)
	}
}
```

```yaml
# m2s.yaml
address: :5000
mode: record # record|mock
recorder_log_path: /tmp/mock # path output that records
```
