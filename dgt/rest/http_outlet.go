package rest

import (
	"encoding/json"
	"io"
	"net/http"
)

func NewRestDataSource(baseUrl string) *DataSource {
	return &DataSource{baseUrl}
}

type DataSource struct {
	BaseUrl string
}

type RequestWrapper struct {
	ref *http.Request
}

type ResponseWrapper struct {
	Status string // e.g. "200 OK"
	Header http.Header
	Body   string
}

func (ds DataSource) NewRequest(schema string) (*RequestWrapper, error) {
	request := new(struct {
		Method string
		Url    string
		Body   string
	})
	json.Unmarshal([]byte(schema), request)
	httpRequest, err := http.NewRequest(request.Method, request.Url, nil)
	if err == nil {
		return &RequestWrapper{ref: httpRequest}, nil
	} else {
		return nil, err
	}
}

func (r RequestWrapper) Apply() (string, error) {
	resp, err := http.DefaultClient.Do(r.ref)
	if err == nil {
		data, _ := json.Marshal(wrapHttpResponse(resp))
		return string(data), nil
	} else {
		return "", err
	}
}

func wrapHttpResponse(resp *http.Response) *ResponseWrapper {
	body, _ := io.ReadAll(resp.Body)
	return &ResponseWrapper{
		Status: resp.Status,
		Header: resp.Header,
		Body:   string(body),
	}
}
