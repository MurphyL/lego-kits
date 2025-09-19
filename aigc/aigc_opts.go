package aigc

import "net/http"

func ServiceProvider(url, token string) AgentOption {
	return func(opts *AgentOptions) {
		opts.ServiceUrl = url
		opts.Token = token
	}
}

func CompletionRequestHook(completionPreHook func(r *http.Request)) AgentOption {
	return func(opts *AgentOptions) {
		opts.CompletionPreHook = completionPreHook
	}
}
