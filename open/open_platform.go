package open

import (
	"net/http"

	"github.com/MurphyL/lego-kits/open/internal/platform"
)

func NewPlatformApp(key, secret string, withOptions ...func(*platform.App)) PlatformApp {
	return platform.NewApp(key, secret, withOptions...)
}

func NewPlatformError(code, phrase, desc string) PlatformError {
	return platform.NewStatus(code, phrase, desc)
}

func NewResultWithCode[T any](code uint, message string, payload T) PlatformResult[T] {
	return platform.NewResultWithCode[T](code, message, payload)
}

func NewPagingWithCode[T any](code uint, total uint, records []T) PlatformResult[*platform.PagingPayload[T]] {
	return platform.NewResultWithCode[*platform.PagingPayload[T]](code, "OK", platform.NewPagingPayload[T](total, records))
}

func WithHttpClient(httpClient *http.Client) func(*platform.App) {
	return func(app *platform.App) {
		app.HttpClient = httpClient
	}
}

func WithHttpRequestBuilder(builder func(*http.Request)) func(*platform.App) {
	return func(app *platform.App) {
		app.RequestBuilder = builder
	}
}

type PlatformApp interface {
	DoRequest(r *http.Request) (*http.Response, error)
	AppKey() string
	AppSecret() string
}

type PlatformError interface {
	Code() string
	Phrase() string
	Desc() string
}

type PlatformAssistant interface {
	PlatformName() string
	PlatformSite() string
}

type PlatformResult[T any] interface {
	Code() uint
	Message() string
	Payload() T
}
