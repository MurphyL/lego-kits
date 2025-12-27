package open

import (
	"github.com/MurphyL/lego-kits/open/internal/platform"
	"github.com/MurphyL/lego-kits/open/internal/result"
	"net/http"
)

func NewPlatformApp(key, secret string) PlatformApp {
	return platform.NewApp(key, secret)
}

func NewPlatformError(code, phrase, desc string) PlatformError {
	return platform.NewStatusCode(code, phrase, desc)
}

func NewResultWithCode[T any](code uint, message string, payload T) PlatformResult[T] {
	return result.NewResultWithCode[T](code, message, payload)
}

func NewPagingWithCode[T any](code uint, total uint, records []T) PlatformResult[*result.PagingPayload[T]] {
	return result.NewResultWithCode[*result.PagingPayload[T]](code, "OK", result.NewPagingPayload[T](total, records))
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
