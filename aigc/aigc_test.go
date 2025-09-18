package aigc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestOllama(t *testing.T) {
	agent := NewAgent(
		WithServiceProvider("http://localhost:11434/api/chat", "qwen3:8b"),
	)
	functionsDictData, err := os.ReadFile("./test_functions.json")
	if nil != err {
		t.Error("读取自定义函数出错：", err.Error())
	}
	var functions []map[string]any
	json.Unmarshal(functionsDictData, &functions)
	messages := []CompletionMessage{{Role: User, Content: "把武汉的天气发送到企微群里"}}
	resp, err := agent.ApplyChatCompletion("qwen3:8b", messages)
	if nil != err {
		t.Error("调用 Ollama 失败", err.Error())
	} else {
		bytes, _ := io.ReadAll(resp.Body)
		t.Log(string(bytes))
	}
}

func TestOllamaV1(t *testing.T) {
	agent := NewAgent(
		WithServiceProvider("http://localhost:11434/api/chat", "qwen3:8b"),
	)
	functionsDictData, err := os.ReadFile("./test_functions.json")
	if nil != err {
		t.Error("读取自定义函数出错：", err.Error())
	}
	var functions []map[string]any
	json.Unmarshal(functionsDictData, &functions)
	messages := []CompletionMessage{{Role: User, Content: "整理2025年内的假日和补班日，以如下JSON结构返回数据：[{\"dt\":\"日期\", \"kind\":\"补板或休假\", \"reason\":\"假日名称或补办原因\"}]"}}
	resp, err := agent.ApplyChatCompletion("qwen3:8b", messages)
	if nil != err {
		t.Error("调用 Ollama 失败", err.Error())
	} else {
		bytes, _ := io.ReadAll(resp.Body)
		t.Log(string(bytes))
	}
}

func TestXfyunX1(t *testing.T) {
	token, ok := os.LookupEnv("XFYUNAI_X1_TOKEN")
	agent := NewAgent(
		WithServiceProvider("https://spark-api-open.xf-yun.com/v2/chat/completions", token),
		WithCompletionRequestHook(func(request *http.Request) {
			request.Header.Add("ContentItem-Type", "application/json")
			request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		}),
	)
	if !ok {
		t.Error("未配置讯飞云 X1 的 Token")
	}
	messages := []CompletionMessage{{Role: User, Content: "武汉今天的天气怎么样？"}}
	resp, err := agent.ApplyChatCompletion("x1", messages)
	if nil != err {
		t.Error("调用 Xfyun AI 失败", err.Error())
	} else {
		bytes, _ := io.ReadAll(resp.Body)
		t.Log(string(bytes))
	}
}

func TestName(t *testing.T) {
	msg := CompletionRequest{ToolChoice: Auto}
	t.Log(msg)
}
