package open

import (
	"github.com/MurphyL/lego-kits/open/internal/location"
)

func GetPublicLocation() (Location, bool) {
	return location.GetPublicLocation()
}

type Location interface {
	GetIP() string
	GetPlace() string
	GetISP() string
}
