package open

import (
	"github.com/MurphyL/lego-kits/open/internal/location"
)

func GetPublicLocation() PlatformResult[Location] {
	if ret, ok := location.GetPublicLocation(); ok {
		return NewResultWithCode[Location](0, "操作成功", ret)
	} else {
		return NewResultWithCode[Location](1, "未知错误", nil)
	}
}

type Location interface {
	IP() string
	Place() string
	ISP() string
}
