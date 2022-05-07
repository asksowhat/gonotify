package gonotify

import (
	"fmt"
	"testing"
)

func TestSendText(t *testing.T) {
	qywx := &QywxClient{
		CorpId:     "ww7fd2960331f0033",
		CorpSecret: "5DryRrtLihp3gqd9O0QNpgTwZqPXaQaIajbSfO1trY",
		AgentId:    "100002",
	}
	// 获取accessToken
	qywx.GetAccessToken()

	messageList := []Message{
		qywx,
	}
	result := false
	for _, s := range messageList {
		result = s.SendText("消息提醒234")
	}
	fmt.Println(result)
}
