package platform

import "net/http"

func NewApp(key, secret string, withOptions ...func(*App)) *App {
	app := &App{key: key, secret: secret}
	for _, withOption := range withOptions {
		withOption(app)
	}
	return app
}

type App struct {
	key, secret    string // 开放平台基础信息
	HttpClient     *http.Client
	RequestBuilder func(*http.Request)
}

func (a *App) AppKey() string {
	return a.key
}

func (a *App) AppSecret() string {
	return a.secret
}

func (a *App) DoRequest(r *http.Request) (*http.Response, error) {
	if a.RequestBuilder != nil {
		a.RequestBuilder(r)
	}
	if nil == a.HttpClient {
		return http.DefaultClient.Do(r)
	} else {
		return a.HttpClient.Do(r)
	}
}
