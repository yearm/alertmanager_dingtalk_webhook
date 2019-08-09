package alert

import (
	"alertmanager_dingtalk_webhook/types"
	"bytes"
	"github.com/json-iterator/go"
	"log"
	"net/http"
)

func Send(alertMessage types.AlertMessage, robotUrl string) error {
	dingTalkMsg := types.AlertToDingTalk(alertMessage)
	data, err1 := jsoniter.Marshal(dingTalkMsg)
	if err1 != nil {
		return err1
	}
	client := &http.Client{}
	resp, err2 := client.Post(robotUrl, "application/json", bytes.NewBuffer(data))
	defer resp.Body.Close()
	if err2 != nil {
		return err2
	}
	log.Printf("Response Status: %s, Header: %s", resp.Status, resp.Header)
	return nil
}
