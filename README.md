# `go-kubernetes-testing`

`go-kubernetes-testing` provides utilities for Kubernetes testing.

## Documentation

Official `godoc` documentation (with examples) can be found at the [Package Registry](https://pkg.go.dev/github.com/x-ethr/go-kubernetes-testing).

## Usage

###### Add Package Dependency

```bash
go get -u github.com/x-ethr/go-kubernetes-testing
```

###### Import & Implement

`main.go`

```go
package main

import (
	"context"
	"fmt"

	"github.com/x-ethr/go-kubernetes-testing"
)

func main() {
	ctx := context.Background()

	instance := secrets.New()
	if e := instance.Walk(ctx, "/etc/secrets"); e != nil {
		panic(e)
	}

	for secret, keys := range instance {
		for key, value := range keys {
			fmt.Println("Secret", secret, "Key", key, "Value", value)
		}
	}

	service := instance["service"]

	port := service["port"]
	hostname := service["hostname"]
	username := service["username"]
	password := service["password"]

	fmt.Println("Port", port, "Hostname", hostname, "Username", username, "Password", password)
}

```

- Please refer to the [code examples](./example_test.go) for additional usage and implementation details.
- See https://pkg.go.dev/github.com/x-ethr/go-kubernetes-testing for additional documentation.

## Contributions

See the [**Contributing Guide**](./CONTRIBUTING.md) for additional details on getting started.
