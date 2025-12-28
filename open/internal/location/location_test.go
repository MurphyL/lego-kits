package location

import (
	"testing"
)

func TestGetPublicLocation(t *testing.T) {
	t.Log(GetPublicLocation())
}

func TestGetPublicIP(t *testing.T) {
	t.Log(GetPublicIP())
}
