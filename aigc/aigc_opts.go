package aigc

import "net/http"

func WithServiceProvider(url, token string) AgentOption {
	return func(opts *AgentOptions) {
		opts.ServiceUrl = url
		opts.Token = token
	}
}

func WithCompletionRequestHook(completionPreHook func(r *http.Request)) AgentOption {
	return func(opts *AgentOptions) {
		opts.CompletionPreHook = completionPreHook
	}
}
