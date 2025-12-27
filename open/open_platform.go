package open

import (
	"github.com/MurphyL/lego-kits/open/internal/platform"
	"net/http"
)

func NewPlatformApp(key, secret string) PlatformApp {
	return platform.NewApp(key, secret)
}

type PlatformApp interface {
	ApplyEndpoint(r *http.Request) (*http.Response, error)
	AppKey() string
	AppSecret() string
}

type Assistant interface {
	PlatformName() string
	PlatformSite() string
}
