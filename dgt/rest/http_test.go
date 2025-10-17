package rest

import (
	"testing"
)

func TestHttpRequest(t *testing.T) {
	ds := DataSource{}
	request, _ := ds.NewRequest(RequestSchema{Url: "https://baidu.com"})
	ret, err := request.Apply()
	if err == nil && ret.Success() {
		t.Log(ret.Attrs())
	}
}
