package aigc

import (
	"encoding/json"
	"io"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	agent := NewCompletionAgent(
		WithServiceProvider("http://localhost:11434/api/chat", "qwen3:8b"),
	)
	functionsDictData, err := os.ReadFile("./test_functions.json")
	if nil != err {
		t.Error("读取自定义函数出错：", err.Error())
	}
	var functions []map[string]any
	json.Unmarshal(functionsDictData, &functions)
	message := map[string]any{
		"model":       "qwen3:8b",
		"messages":    []any{map[string]string{"role": "user", "content": "把武汉的天气发送到企微群里"}},
		"stream":      false,
		"tools":       functions,
		"tool_choice": "required",
	}
	resp, err := agent.ApplyApplyCompletion(message)
	if nil != err {
		t.Error("调用 LLM 失败", err.Error())
	} else {
		bytes, _ := io.ReadAll(resp.Body)
		t.Log(string(bytes))
	}
}
