package aigc

type CompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func WithMessage(role, content string) func(*ChatCompletion) {
	return func(request *ChatCompletion) {
		request.Messages = append(request.Messages, CompletionMessage{Role: role, Content: content})
	}
}

func WithUserMessage(content string) func(*ChatCompletion) {
	return WithMessage("user", content)
}

func WithSystemMessage(content string) func(*ChatCompletion) {
	return WithMessage("system", content)
}

func WithAssistantMessage(content string) func(*ChatCompletion) {
	return WithMessage("assistant", content)
}
