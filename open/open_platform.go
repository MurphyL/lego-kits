package open

import (
	"github.com/MurphyL/lego-kits/open/internal/platform"
	"net/http"
)

func NewPlatformApp(key, secret string) PlatformApp {
	return &platform.App{Key: key, Secret: secret}
}

type PlatformApp interface {
	ApplyEndpoint(r *http.Request) (*http.Response, error)
}

type Assistant interface {
	PlatformName() string
	PlatformSite() string
}
