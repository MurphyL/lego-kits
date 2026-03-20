package wecom

const (
	// 企业API基础URL
	qyapiBaseURL = "https://qyapi.weixin.qq.com/cgi-bin"
)

func NewMessagePushAgent(key string) WebhookAgent {
	return &QyapiWebhookAgent{baseURL: qyapiBaseURL}
}

type WebhookAgent interface {
}

// QyapiWebhookAgent 企业API钩子代理
type QyapiWebhookAgent struct {
	baseURL string
}
