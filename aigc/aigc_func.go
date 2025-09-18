package aigc

type FunctionSchema struct {
}

func IsValidFunction(vars map[string]any) bool {
	if nil != vars {
		var ok bool
		var subVar any
		subVar, ok = vars["type"]
		if ok && subVar == "function" {
			subVar, ok = vars["function"]
			if ok {
				return true
			}
		}
	}
	return false
}
