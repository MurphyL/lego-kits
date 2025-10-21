package aigc

import "net/http"

func CompletionRequestHook(completionPreHook func(r *http.Request)) AgentOption {
	return func(opts *AgentOptions) {
		opts.CompletionPreHook = completionPreHook
	}
}
