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

