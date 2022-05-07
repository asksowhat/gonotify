// 企业微信应用消息
package gonotify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Message interface {
	SendText(text string) bool
	GetAccessToken() *QywxClient
}

type QywxClient struct {
	CorpId      string
	CorpSecret  string
	AccessToken string
	AgentId     string
}

func (qc *QywxClient) SendText(text string) bool {
	params := map[string]interface{}{
		"msgtype": "text",
		"touser":  "@all",
		"agentid": qc.AgentId,
		"text": map[string]interface{}{
			"content": text,
		},
	}
	postApi(qc, "https://qyapi.weixin.qq.com/cgi-bin/message/send", params)
	return true
}

func (qc *QywxClient) GetAccessToken() *QywxClient {
	urlValues := url.Values{
		"corpid":     {qc.CorpId},
		"corpsecret": {qc.CorpSecret},
	}
	resp, err := http.Get("https://qyapi.weixin.qq.com/cgi-bin/gettoken?" + urlValues.Encode())
	if err != nil {
		fmt.Println(err)
		return &QywxClient{}
	}
	body, _ := ioutil.ReadAll(resp.Body)
	strBody := string(body)
	resultMap := make(map[string]interface{}, 10)
	json.Unmarshal([]byte(strBody), &resultMap)
	accessToken, ok := resultMap["access_token"]
	if !ok {
		fmt.Println("获取accetoken失败")
		return &QywxClient{}
	}
	qc.AccessToken = accessToken.(string)
	return qc
}

func postApi(qc *QywxClient, url string, urlValues map[string]interface{}) map[string]interface{} {
	accessToken := qc.AccessToken
	url = url + "?access_token=" + accessToken
	jsonValue, _ := json.Marshal(urlValues)
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	body, _ := ioutil.ReadAll(resp.Body)
	strBody := string(body)
	resultMap := make(map[string]interface{}, 10)
	fmt.Println(strBody)
	json.Unmarshal([]byte(strBody), &resultMap)
	return resultMap
}
