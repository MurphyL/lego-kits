package aigc

// ToolChoiceMode 可选值为string或json对象类型，可以控制模型调用哪些函数（如果有）。可使用 {"type": "function", "function": {"name": "my_function"}} 来强制模型指定函数。
type ToolChoiceMode string

const (
	ToolChoiceModeAuto     ToolChoiceMode = "auto" // 模型可以在生成消息或调用函数之间进行选择。
	ToolChoiceModeNone     ToolChoiceMode = "none" // 模型不会调用函数而是生成消息。
	ToolChoiceModeRequired ToolChoiceMode = "required"
)

func WithFunction(name, desc string, params map[string]any) func(*ChatCompletion) {
	return func(request *ChatCompletion) {
		functionBody := map[string]any{"name": name, "description": desc, "parameters": params}
		request.Tools = append(request.Tools, map[string]any{"type": "function", "function": functionBody})
	}
}
