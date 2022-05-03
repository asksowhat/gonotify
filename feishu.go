// 飞书自定义机器人
package gonotify

import (
	"bytes"
	"fmt"
	"time"

	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Message interface {
	SendText(text string) bool
	SendRichText(title string, content []map[string]interface{}) bool
}

type FeishuClient struct {
	Url    string
	Secret string
}

// 发送文本消息
// 参考文档 https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
func (fc *FeishuClient) SendText(text string) bool {
	params := map[string]interface{}{
		"msg_type": "text",
		"content": map[string]interface{}{
			"text": text,
		},
	}
	result := PostApi(fc, params)
	fmt.Println(result)
	if _, ok := result["code"]; ok {
		return false
	}
	return true
}

// 发送富文本消息
// 参考文档 https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
func (fc *FeishuClient) SendRichText(title string, content []map[string]interface{}) bool {
	params := map[string]interface{}{
		"msg_type": "post",
		"content": map[string]interface{}{
			"post": map[string]interface{}{
				"zh_cn": map[string]interface{}{
					"title": title,
					"content": []interface{}{
						content,
					},
				},
			},
		},
	}
	result := PostApi(fc, params)
	fmt.Println(result)
	if _, ok := result["code"]; ok {
		return false
	}
	return true
}

// 签名算法
func GenSign(secret string, timestamp int64) (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret

	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}

// post请求
func PostApi(fc *FeishuClient, urlValues map[string]interface{}) map[string]interface{} {
	timestamp := time.Now().Unix()
	if len(fc.Secret) != 0 {
		urlValues["timestamp"] = timestamp
		urlValues["sign"], _ = GenSign(fc.Secret, timestamp)
	}
	jsonValue, _ := json.Marshal(urlValues)
	resp, _ := http.Post(fc.Url, "application/json", bytes.NewBuffer(jsonValue))
	body, _ := ioutil.ReadAll(resp.Body)
	strBody := string(body)
	resultMap := make(map[string]interface{}, 10)
	json.Unmarshal([]byte(strBody), &resultMap)
	return resultMap
}
