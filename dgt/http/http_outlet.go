package http

import (
	"io"
	"net/http"
)

type Request struct {
	Method string
	Url    string
	Header http.Header
	Body   io.Reader
}

func (r Request) Exec() (string, error) {
	var err error
	var ret string
	httpRequest, err := http.NewRequest(r.Method, r.Url, r.Body)
	if err == nil {
		if nil != r.Header {
			httpRequest.Header = r.Header
		}
		resp, err := http.DefaultClient.Do(httpRequest)
		if nil == err {
			data, err := io.ReadAll(resp.Body)
			if err == nil {
				ret = string(data)
			}
		}
	}
	return ret, err
}
