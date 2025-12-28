package location

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// 开放IP地址获取的接口
var publicAddrAPIS = []string{
	"https://api.ipify.org",
	"https://icanhazip.com/",
	"http://checkip.amazonaws.com",
	"https://myexternalip.com/raw",
	"https://www.trackip.net/IP",
	"https://ipecho.net/plain",
}

type IPLocation struct {
	IP    string
	Place string
	ISP   string
}

func (l *IPLocation) GetIP() string {
	return l.IP
}

func (l *IPLocation) GetPlace() string {
	return l.Place
}

func (l *IPLocation) GetISP() string {
	return l.ISP
}

func GetPublicLocation() (*IPLocation, bool) {
	var err error
	var resp *http.Response
	if resp, err = http.Get("https://myip.ipip.net/"); err == nil {
		data, _ := io.ReadAll(resp.Body)
		parts := strings.SplitN(string(data), "  ", 3)
		return &IPLocation{
			IP:    strings.TrimPrefix(parts[0], "当前 GetIP："),
			Place: strings.TrimPrefix(parts[1], "来自于："),
			ISP:   parts[2],
		}, true
	}
	return nil, false
}

// GetPublicIP 获取公网 GetIP
func GetPublicIP() (string, bool) {
	var err error
	var resp *http.Response
	// 直接返回 GetIP 的接口
	for _, api := range publicAddrAPIS {
		if resp, err = http.Get(api); err == nil {
			data, _ := io.ReadAll(resp.Body)
			return string(data), true
		}
	}
	// 服用 ipip.net 的接口
	if ret, ok := GetPublicLocation(); ok {
		return ret.IP, true
	}
	// 需要使用 JSON 解析数据的接口
	if resp, err = http.Get("https://openapi.lddgo.net/base/gtool/api/v1/GetIp"); err == nil {
		data, _ := io.ReadAll(resp.Body)
		var ret map[string]any
		if err = json.Unmarshal(data, &ret); err == nil && ret["code"] == 0 && nil != ret["data"] {
			return ret["data"].(map[string]string)["ipv4"], true
		}
	}
	return "", false
}
