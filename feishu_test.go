package gonotice

import (
	"testing"
)

func TestSendText(t *testing.T) {
	feishu := &FeishuClient{
		Url:    "https://open.feishu.cn/open-apis/bot/v2/hook/62935edf-28ca-4136-84f5-03415153eeb", // webhook地址
		Secret: "QofaUxbZEHakjshdasjkdaks",                                                         // 可选 加密参数
	}
	messageList := []Message{
		feishu,
	}
	result := false
	for _, s := range messageList {
		result = s.SendText("消息提醒234")
	}
	if !result {
		t.Error("错误")
	}
}

func TestSendRichText(t *testing.T) {
	feishu := &FeishuClient{
		Url:    "https://open.feishu.cn/open-apis/bot/v2/hook/62935edf-28ca-4136-84f5-03415153eeb", // webhook地址
		Secret: "QofaUxbZEHakjshdasjkdaks",                                                         // 可选 加密参数
	}
	messageList := []Message{
		feishu,
	}
	result := false
	for _, s := range messageList {
		content := []map[string]interface{}{
			{
				"tag":  "text",
				"text": "项目有更新",
			},
			{
				"tag":  "a",
				"text": "请查看",
				"href": "http://www.example.com",
			},
			{
				"tag":     "at",
				"user_id": "ou_18eac8********17ad4f02e8bbbb",
			},
		}
		result = s.SendRichText("消息提醒", content)
	}
	if !result {
		t.Error("错误")
	}
}
