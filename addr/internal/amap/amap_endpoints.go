package amap

import (
	"log"
	"net/http"
	"net/url"
)

type SearchAroundOption func(*url.Values)

// SearchAround 高德地图 - POI2.0 - 周边搜索 - https://restapi.amap.com/v5/place/around?parameters
func (a *Assistant) SearchAround(keywords string, withOptions ...SearchAroundOption) (any, bool) {
	u, _ := url.ParseRequestURI("https://restapi.amap.com/v5/place/text")
	q := url.Values{"key": []string{a.platformApp.AppKey()}, "keywords": []string{keywords}}
	u.RawQuery = q.Encode()
	if ret, err := a.ApplyRequest(&http.Request{URL: u}, "pois"); err == nil {
		return ret, true
	} else {
		log.Panicln("POI2.0 接口出错：", err)
		return nil, false
	}
}

// ReGEO 高德地图 - 地理/逆地理编码 - https://lbs.amap.com/api/webservice/guide/api/georegeo
func (a *Assistant) ReGEO(address string, withOptions ...SearchAroundOption) (any, bool) {
	u, _ := url.ParseRequestURI("https://restapi.amap.com/v3/geocode/geo")
	q := url.Values{"key": []string{a.platformApp.AppKey()}, "address": []string{address}}
	u.RawQuery = q.Encode()
	if ret, err := a.ApplyRequest(&http.Request{URL: u}, "geocodes"); err == nil {
		return ret, true
	} else {
		log.Panicln("POI2.0 接口出错：", err)
		return nil, false
	}
}
