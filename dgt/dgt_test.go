package dgt

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func wrapHttpResponse(resp *http.Response) *ResponseWrapper {
	body, _ := io.ReadAll(resp.Body)
	return &ResponseWrapper{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       string(body),
	}
}

type ResponseWrapper struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Header     http.Header
	Body       string
}

func TestHttpRequest(t *testing.T) {
	resp, err := http.Get("https://baidu.com")
	if err == nil {
		ret := wrapHttpResponse(resp)
		data, _ := json.Marshal(ret)
		t.Log(string(data))
	}
}
