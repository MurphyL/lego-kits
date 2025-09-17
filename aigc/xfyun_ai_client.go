package aigc

type XfyunAiClient struct {
	Model string
	Token string
}

/**
func (client XfyunAiClient) ApplyCompletion(text string, opts *CompletionOptions) CompletionResponse {
	var baseURL string
	switch client.Model {
	case "x1":
		// X1 模型文档 - https://www.xfyun.cn/doc/spark/X1http.html
		baseURL = "https://spark-api-open.xf-yun.com/v2/chat/completions"
	}
	content, _ := GetCompletionRequestMessage(text, client.Model, opts.Streaming, opts.WebSearch)
	request, _ := http.NewRequest("POST", baseURL, bytes.NewReader(content))
	request.Header.Add("ContentItem-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token))
	resp, err := http.DefaultClient.Do(request)
	if nil == err {
		return XfyunCompletionResponse{resp: resp}
	} else {
		return XfyunCompletionResponse{}
	}
}

type XfyunCompletionResponse struct {
	resp *http.Response
}

func (x XfyunCompletionResponse) IsValid() bool {
	return nil != x.resp && x.resp.StatusCode == 200
}

func (x XfyunCompletionResponse) GetResult() string {
	bytes, _ := io.ReadAll(x.resp.Body)
	return string(bytes)
}
**/
