package aigc

import (
	"encoding/json"
	"os"
	"testing"
)

func TestOllama(t *testing.T) {
	agent := NewAgent("http://localhost:11434/api/chat", "qwen3:8b")
	functionsDictData, err := os.ReadFile("./test_functions.json")
	if nil != err {
		t.Error("读取自定义函数出错：", err.Error())
	}
	var functions []map[string]any
	json.Unmarshal(functionsDictData, &functions)
	resp, err := agent.ApplyChatCompletion("qwen3:8b", WithUserMessage("把武汉的天气发送到企微群里"))
	if nil != err {
		t.Error("调用 Ollama 失败", err.Error())
	} else {
		bytes, _ := json.Marshal(resp)
		t.Log(string(bytes))
	}
}

func TestOllamaV1(t *testing.T) {
	agent := NewAgent("http://localhost:11434/api/chat", "qwen3:8b")
	functionsDictData, err := os.ReadFile("./test_functions.json")
	if nil != err {
		t.Error("读取自定义函数出错：", err.Error())
	}
	var functions []map[string]any
	json.Unmarshal(functionsDictData, &functions)
	resp, err := agent.ApplyChatCompletion("qwen3:8b", WithUserMessage("整理2025年内的假日和补班日，以如下JSON结构返回数据：[{\"dt\":\"日期\", \"kind\":\"补板或休假\", \"reason\":\"假日名称或补办原因\"}]"))
	if nil != err {
		t.Error("调用 Ollama 失败", err.Error())
	} else {
		bytes, _ := json.Marshal(resp)
		t.Log(string(bytes))
	}
}

func TestXfyunX1(t *testing.T) {
	token, ok := os.LookupEnv("XFYUN_AI_X1_TOKEN")
	if ok {
		agent := NewAgent("https://spark-api-open.xf-yun.com/v2/chat/completions", token)
		resp, err := agent.ApplyChatCompletion("x1", WithUserMessage("武汉今天的天气怎么样？"))
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

func TestName(t *testing.T) {
	msg := CompletionRequest{ToolChoice: ToolChoiceModeAuto}
	t.Log(msg)
}
