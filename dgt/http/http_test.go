package http

import (
	"testing"
)

func TestHttpRequest(t *testing.T) {
	req := Request{Method: "GET", Url: "https://www.baidu.com"}
	ret, err := req.Exec()
	if err == nil {
		t.Log("HTTP 数据请求成功：", ret)
	} else {
		t.Log("HTTP 数据请求出错：", err.Error())
	}
}
