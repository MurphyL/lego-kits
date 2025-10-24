package misc

import (
	"runtime"
	"strings"
)

// IsTesting 判断是否在单元测试中运行
func IsTesting() bool {
	// 遍历调用栈
	for i := 0; ; i++ {
		pc, _, _, ok := runtime.Caller(i)
		if !ok {
			break // 遍历完调用栈仍未找到测试相关函数
		}
		// 获取当前调用栈帧的函数信息
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		funcName := fn.Name()
		// 检查函数名是否包含测试相关标识：
		// 1. 来自testing包（如testing.T.Run）
		// 2. 测试函数（以Test开头）
		if strings.HasPrefix(funcName, "testing.") || strings.HasPrefix(funcName, "Test") {
			return true
		}
	}
	return false
}
