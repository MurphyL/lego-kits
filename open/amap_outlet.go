package open

import (
	"net/url"

	"github.com/MurphyL/lego-kits/open/internal/amap"
)

func NewAmapClient(key, secret string) AmapAssistant {
	return amap.NewAmapClient(key, secret)
}

type AmapAssistant interface {
	SearchAround(keywords string, withOptions ...func(*url.Values)) (any, bool)
	ReGEO(address string, withOptions ...func(*url.Values)) (any, bool)
}

// WithAmapEndpointUrlParams 高德地图 - 添加动态参数
func WithAmapEndpointUrlParams(key, value string) func(*url.Values) {
	return func(values *url.Values) {
		values.Add(key, value)
	}
}
