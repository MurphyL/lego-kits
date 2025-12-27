package amap

import (
	"encoding/json"
	"fmt"
	"github.com/MurphyL/lego-kits/open"
	"io"
	"net/http"
	"net/url"
)

func NewAmapClient(key, secret string) *Assistant {
	platformApp := open.NewPlatformApp(key, secret)
	endpointRoot, _ := url.Parse("https://restapi.amap.com")
	return &Assistant{platformApp: platformApp, endpointRoot: endpointRoot}
}

type Assistant struct {
	endpointRoot *url.URL
	platformApp  open.PlatformApp
}

func (a *Assistant) PlatformName() string {
	return "高德地图/服务端应用"
}

func (a *Assistant) PlatformSite() string {
	return "https://lbs.amap.com/api"
}

func (a *Assistant) ApplyRequest(r *http.Request, key string) (any, error) {
	var err error
	if resp, e1 := a.platformApp.DoRequest(r); e1 == nil {
		data, _ := io.ReadAll(resp.Body)
		ret := ParsedAmapResult{}
		if e2 := json.Unmarshal(data, &ret); e2 == nil {
			if ret.Success() {
				return ret.Get(key), nil
			}
			err = fmt.Errorf("获取高德地图开放平台返回的结果出错：%v", ret["info"])
		} else {
			err = fmt.Errorf("解析高德地图开放平台返回的结果出错：%v", e2)
		}
	} else {
		err = fmt.Errorf("调用高德地图开放平台接口出错：%v", e1.Error())
	}
	return nil, err
}
