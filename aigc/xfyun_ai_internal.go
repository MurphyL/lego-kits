package aigc

import "encoding/json"

/* 讯飞云相关工具类 */

func GetCompletionRequestMessage(text string, model string, streaming bool, enableWebSearch bool) ([]byte, error) {
	payload := map[string]any{
		"model":  model,
		"stream": streaming,
		"messages": []map[string]any{
			{"role": "user", "content": text},
		},
	}
	if enableWebSearch {
		payload["tools"] = []map[string]any{
			{
				"type": "web_search",
				"web_search": map[string]any{
					"enable":      enableWebSearch,
					"search_mode": "deep",
				},
			},
		}
	}
	return json.Marshal(payload)
}
