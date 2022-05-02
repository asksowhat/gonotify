# 背景

本项目始于学习golang，目的在于简化多个项目间给第三方发送通知的场景。预计会对接飞书、钉钉、企业微信、telegram等平台。

## 飞书

测试前，需先准备好飞书自定义机器人webhook地址

### 发送文本

```bash
go test -v feishu_test.go feishu.go -test.run TestSendText
```

### 发送富文本

```bash
go test -v feishu_test.go feishu.go -test.run TestSendRichText
```