package addr

import (
	"github.com/MurphyL/lego-kits/addr/internal/amap"
	"net/url"
)

func NewAmapClient(key, secret string) *amap.Assistant {
	return amap.NewAmapClient(key, secret)
}

// WithAmapEndpointUrlParams 高德地图 - 添加动态参数
func WithAmapEndpointUrlParams(key, value string) amap.SearchAroundOption {
	return func(values *url.Values) {
		values.Set(key, value)
	}
}
