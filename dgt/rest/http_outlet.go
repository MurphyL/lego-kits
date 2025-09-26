package rest

import (
	"io"
	"net/http"
	"strings"

	"github.com/MurphyL/lego-kits/dgt"
)

type DataSource struct {
}

type RequestWrapper struct {
	*http.Request
}

type ResponseWrapper struct {
	*http.Response
}

type RequestSchema struct {
	Method string
	Url    string
	Body   string
}

func (ds DataSource) NewRequest(request RequestSchema) (*RequestWrapper, error) {
	req, err := http.NewRequest(
		request.Method,
		request.Url,
		strings.NewReader(request.Body),
	)
	return &RequestWrapper{Request: req}, err
}

func (r RequestWrapper) Apply() (dgt.Response, error) {
	resp, err := http.DefaultClient.Do(r.Request)
	return ResponseWrapper{Response: resp}, err
}

func (resp ResponseWrapper) Success() bool {
	return resp.StatusCode == 200
}

func (resp ResponseWrapper) Attrs() map[string]string {
	vars := make(map[string]string)
	for k := range resp.Header {
		vars[k] = strings.Join(resp.Header.Values(k), ",")
	}
	return vars
}

func (resp ResponseWrapper) Attr(key string) string {
	return resp.Header.Get(key)
}

func (resp ResponseWrapper) Body() string {
	data, _ := io.ReadAll(resp.Response.Body)
	return string(data)
}
