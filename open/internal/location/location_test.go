package location

import (
	"strings"
	"testing"
)

func TestGetPublicLocation(t *testing.T) {
	t.Log(GetPublicLocation())
}

func TestGetPublicIP(t *testing.T) {
	t.Log(GetPublicIP())
}

func TestGet(t *testing.T) {
	_, IP, _ := strings.Cut("当前 IP：59.175.123.149", "：")
	t.Log(IP)
}
