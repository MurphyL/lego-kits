package open

import (
	"github.com/MurphyL/lego-kits/open/internal/location"
)

func GetPublicLocation() PlatformResult[*location.IPLocation] {
	if ret, ok := location.GetPublicLocation(); ok {
		return NewResultWithCode(0, "操作成功", ret)
	} else {
		return NewResultWithCode[*location.IPLocation](1, "未知错误", nil)
	}
}
