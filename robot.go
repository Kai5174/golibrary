package golibrary

import (
	"bytes"
	"fmt"
	"github.com/blinkbean/dingtalk"
	"net/http"
)

// SendDingTalk 向钉钉机器人发送消息
func SendDingTalk(dingToken []string, message string) error {
	cli := dingtalk.InitDingTalk(dingToken, ".")
	err := cli.SendTextMessage(message)
	if err != nil {
		return err
	}
	return nil
}

// SendWechat 向微信机器人发送消息
func SendWechat(wechatUrl, message string) error {
	client := &http.Client{}
	jsonData := []byte(fmt.Sprintf("{\r\n  \"msgtype\": \"text\",\r\n  \"text\": {\r\n    \"content\": \"%s\"\r\n  }\r\n}", message))
	request, err := http.NewRequest("POST", wechatUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	_, err = client.Do(request)
	return err
}
