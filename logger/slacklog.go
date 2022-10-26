package logger

import (
	"encoding/json"
	"go-log-slack/goslack"
)

var goSlackClient *goslack.GoSlackClient
var serviceName string

func SetSlackLogger(webhookUrl, service string) {
	client := goslack.NewGoSlackClient(webhookUrl)
	goSlackClient = &client
	serviceName = service
}

func send(msg string) {
	clientReq := goslack.ClientRequest{
		Header:      "Alert",
		ServiceName: serviceName,
		Summary:     "Some error occurred TEST",
		Details:     msg,
		Status:      goslack.Alert,
	}

	_ = goSlackClient.Send(clientReq)
}

func ProcessAndSend(slackLogReq SlackLogRequest, status int, logType string) error{

	if goSlackClient != nil {
		msg, err := json.MarshalIndent(&slackLogReq, "", "\t")
		if err != nil {
			return err
		}
		if msg != nil {
			clientReq := goslack.ClientRequest{
				Header:      slackLogReq.Level,
				ServiceName: serviceName,
				Summary:     logType + " Log from " + serviceName,
				Details:     string(msg),
				Status:      status,
			}
			err = goSlackClient.Send(clientReq)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
