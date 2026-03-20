package open

import (
	"murphyl.com/lego/oapi/internal/location"
)

func GetPublicLocation() (Location, bool) {
	return location.GetPublicLocation()
}

type Location interface {
	GetIP() string
	GetPlace() string
	GetISP() string
}
