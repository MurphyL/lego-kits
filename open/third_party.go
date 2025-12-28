package open

import (
	"net/http"

	"github.com/MurphyL/lego-kits/open/internal/third_party"
)

/* 开放平台 - 三方服务的接口封装 */

func NewThirdPartyApp(key, secret string, withOptions ...func(*third_party.App)) ThirdPartyApp {
	return third_party.NewApp(key, secret, withOptions...)
}

func NewThirdPartyError(code, phrase, desc string) ThirdPartyError {
	return third_party.NewStatus(code, phrase, desc)
}

func WithHttpClient(httpClient *http.Client) func(*third_party.App) {
	return func(app *third_party.App) {
		app.HttpClient = httpClient
	}
}

func WithHttpRequestBuilder(builder func(*http.Request)) func(*third_party.App) {
	return func(app *third_party.App) {
		app.RequestBuilder = builder
	}
}

type ThirdPartyApp interface {
	DoRequest(r *http.Request) (*http.Response, error)
	AppKey() string
	AppSecret() string
}

type ThirdPartyError interface {
	Code() string
	Phrase() string
	Desc() string
}

type ThirdPartyAssistant interface {
	PlatformName() string
	PlatformSite() string
}
