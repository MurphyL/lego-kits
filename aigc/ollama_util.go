package aigc

func NewOllamaAgent(url string, withOptions ...func(*AgentOptions)) Agent {
	options := &AgentOptions{}
	for _, withOption := range withOptions {
		withOption(options)
	}
	return &ollamaAgent{internalAgent: internalAgent{Url: url, AgentOptions: options}}
}

type ollamaAgent struct {
	internalAgent
}
