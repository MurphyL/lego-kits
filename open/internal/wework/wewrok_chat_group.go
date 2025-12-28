package wework

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/MurphyL/lego-kits/open/internal/third_party"
)

func NewChatGroupPushService(key string, withOptions ...func(*third_party.App)) *ChatGroupPushService {
	webhookURL := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + key
	platformApp := third_party.NewApp(key, "", withOptions...)
	platformApp.RequestBuilder = func(request *http.Request) {
		request.Header.Set("Content-Type", "application/json")
	}
	return &ChatGroupPushService{webhookURL: webhookURL, platformApp: platformApp}
}

type ChatGroupPushService struct {
	webhookURL  string
	platformApp *third_party.App
}

func (c *ChatGroupPushService) PlatformName() string {
	return "企业微信聊天群机器人"
}

func (c *ChatGroupPushService) PlatformSite() string {
	return "https://developer.work.weixin.qq.com/document/path/91770"
}

func (c *ChatGroupPushService) send(body any) {
	payload, _ := json.Marshal(body)
	r := httptest.NewRequest(http.MethodPost, c.webhookURL, bytes.NewReader(payload))
	if resp, err := c.platformApp.DoRequest(r); err == nil {
		data, _ := io.ReadAll(resp.Body)
		log.Println(string(data))
	} else {
		log.Println(err)
	}
}

func (c *ChatGroupPushService) SendTextMessage(text string) {
	body := map[string]any{"msgtype": "text", "text": map[string]any{"content": text}}
	c.send(body)
}
