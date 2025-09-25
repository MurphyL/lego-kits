package rest

import (
	"encoding/json"
	"testing"
)

func TestHttpRequest(t *testing.T) {
	client := NewRestDataSource("https://www.baidu.com")
	schema, _ := json.Marshal(RequestSchema{})
	request, _ := client.NewRequest(string(schema))
	t.Log(request)
}
