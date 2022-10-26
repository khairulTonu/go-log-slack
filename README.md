
# go-log-slack

A go package that will automatically send the messages to the desired slack channel.

## Running Example



```go
package main

import (
	"errors"
	"go-log-slack/logger"
)

func main() {
	webhookUrl := "webhook url"
	service := "service name"
	logger.SetSlackLogger(webhookUrl, service)
	e := errors.New("custom error")
	logger.Error("Error occurred ", e, nil)
}

```

To run tests, run the following command

```bash
  go run main.go
```

