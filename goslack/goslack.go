package goslack

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var client GoSlackClient

type GoSlackClient struct {
	webhookUrl string
}

func NewGoSlackClient(webhookUrl string) GoSlackClient {
	client = GoSlackClient{
		webhookUrl: webhookUrl,
	}
	return client
}

// Send will call api to send a message to the slack channel
func (sc *GoSlackClient) Send(clientReq ClientRequest) error {

	if err := clientReq.Validate(); err != nil {
		return err
	}

	attachments := PrepareAttachmentBody(clientReq)
	slackBody, _ := json.Marshal(SlackRequestBody{Attachments: attachments})
	req, err := http.NewRequest(http.MethodPost, sc.webhookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	if buf.String() != "ok" {
		fmt.Print(resp.Status, "\n")
		return errors.New("non-ok response returned from slack")
	}
	return nil
}
