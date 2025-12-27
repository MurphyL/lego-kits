package open

import "net/http"

func NewPlatformApp(key, secret string) *PlatformApp {
	return &PlatformApp{key, secret}
}

type PlatformApp struct {
	Key, Secret string // 开放平台基础信息
}

func (a *PlatformApp) ApplyEndpoint(r *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(r)
}
