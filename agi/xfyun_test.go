package aigc

import (
	"encoding/json"
	"os"
	"testing"
)

func TestXfyunX1(t *testing.T) {
	if token, ok := os.LookupEnv("XFYUN_AI_X1_TOKEN"); ok {
		agent := NewXfyunAgent("https://spark-api-open.xf-yun.com/v2/chat/completions", token)
		resp, err := agent.ApplyCompletion("x1", WithUserMessage("武汉今天的天气怎么样？"))
		if nil != err {
			t.Error("调用 Xfyun AI 失败", err.Error())
		} else {
			bytes, _ := json.Marshal(resp)
			t.Log(string(bytes))
		}
	} else {
		t.Error("未配置讯飞云 X1 的 Token")
		return
	}
}
