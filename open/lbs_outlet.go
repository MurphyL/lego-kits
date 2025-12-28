package open

import (
	"github.com/MurphyL/lego-kits/open/internal/location"
)

func GetPublicLocation() (*location.IPLocation, bool) {
	return location.GetPublicLocation()
}
