package aigc

func NewXfyunAgent(url, token string, withOptions ...func(*AgentOptions)) SharedAgent {
	options := &AgentOptions{}
	for _, withOption := range withOptions {
		withOption(options)
	}
	return &xfyunAgent{
		internalAgent: internalAgent{Url: url, Token: token, AgentOptions: options},
	}
}

type xfyunAgent struct {
	internalAgent
}
