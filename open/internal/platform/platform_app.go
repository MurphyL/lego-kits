package platform

import "net/http"

func NewApp(key, secret string) *App {
	return &App{key, secret}
}

type App struct {
	key, secret string // 开放平台基础信息
}

func (a *App) AppKey() string {
	return a.key
}

func (a *App) AppSecret() string {
	return a.secret
}

func (a *App) ApplyEndpoint(r *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(r)
}
