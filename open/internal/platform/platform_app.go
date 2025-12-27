package platform

import "net/http"

type App struct {
	Key, Secret string // 开放平台基础信息
}

func (a *App) ApplyEndpoint(r *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(r)
}
