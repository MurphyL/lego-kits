package aigc

type ToolChoiceMode string

const (
	ToolChoiceModeAuto     ToolChoiceMode = "auto"
	ToolChoiceModeNone     ToolChoiceMode = "none"
	ToolChoiceModeRequired ToolChoiceMode = "required"
)

type CompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CompletionRequest struct {
	Messages   []CompletionMessage `json:"messages"`
	Model      string              `json:"model"`
	Streaming  bool                `json:"stream"`
	ToolChoice ToolChoiceMode      `json:"tool_choice,omitempty"`
	Tools      []map[string]any    `json:"tools,omitempty"`
}

type CompletionRequestOption func(*CompletionRequest)

func NewCompletionMessage(role string, content string) CompletionMessage {
	return CompletionMessage{Role: role, Content: content}
}

func NewCompletionTool(kind string, body map[string]any) map[string]any {
	return map[string]any{"type": kind, kind: body}
}

func WithUserMessage(content string) CompletionRequestOption {
	return func(request *CompletionRequest) {
		request.Messages = append(request.Messages, NewCompletionMessage("user", content))
	}
}

func WithSystemMessage(content string) CompletionRequestOption {
	return func(request *CompletionRequest) {
		request.Messages = append(request.Messages, NewCompletionMessage("system", content))
	}
}

func WithAssistantMessage(content string) CompletionRequestOption {
	return func(request *CompletionRequest) {
		request.Messages = append(request.Messages, NewCompletionMessage("assistant", content))
	}
}

func WithFunction(name, desc string, params map[string]any) CompletionRequestOption {
	return func(request *CompletionRequest) {
		functionBody := map[string]any{"name": name, "description": desc, "parameters": params}
		request.Tools = append(request.Tools, NewCompletionTool("function", functionBody))
	}
}
