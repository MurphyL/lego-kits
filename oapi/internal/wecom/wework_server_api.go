package wecom

import "fmt"

const (
	// 企业API基础URL
	qyapiBaseEndpoint = "https://qyapi.weixin.qq.com/cgi-bin"
)

type WebhookAgent struct {
	url string
}

func NewMessagePushAgent(key string) WebhookAgent {
	return WebhookAgent{url: fmt.Sprintf("%s/webhook/send?key=%s", qyapiBaseEndpoint, key)}
}
